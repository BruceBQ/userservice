package audit

import (
	"context"
	"encoding/json"
	"userservice/app"
	"userservice/clog"
	"userservice/jobs"
	tjobs "userservice/jobs/interfaces"
	"userservice/model"
)

const (
	JobName  = "Audit"
	JobQueue = "Queue:Audit"
)

type Worker struct {
	name      string
	stop      chan bool
	stopped   chan bool
	jobs      chan model.Job
	jobServer *jobs.JobServer
	app       *app.App
}

func init() {
	app.RegisterJobsAuditInterface(func(a *app.App) tjobs.AuditInterface {
		return &AuditJobInterfaceImpl{
			App: a,
		}
	})
}

type AuditJobInterfaceImpl struct {
	App *app.App
}

func (m *AuditJobInterfaceImpl) MakeWorker() model.Worker {
	return &Worker{
		name:      JobName,
		stop:      make(chan bool, 1),
		stopped:   make(chan bool, 1),
		jobs:      make(chan model.Job),
		jobServer: m.App.Srv().Jobs,
		app:       m.App,
	}
}

func (worker *Worker) Run() {
	clog.Debug("Worker started", clog.String("worker", worker.name))
	defer func() {
		clog.Debug("Worker finished!", clog.String("worker", worker.name))
	}()

	for {
		select {
		case <-worker.stop:
			clog.Debug("Worker received stop signal", clog.String("worker", worker.name))
			return
		case job := <-worker.jobs:
			clog.Debug("Worker received a new candidate job", clog.String("worker", worker.name))
			worker.DoJob(&job)
		}
	}
}

func (worker *Worker) Stop() {
	worker.stop <- true
	<-worker.stopped
}

func (worker *Worker) JobChannel() chan<- model.Job {
	return worker.jobs
}

func (worker *Worker) DoJob(job *model.Job) {
	result, err := worker.app.Srv().SessionCache.LPop(context.Background(), JobQueue).Result()
	if err != nil {
		clog.Error("Worker: Failed to get audit logs from cache", clog.String("worker", worker.name), clog.String("error", err.Error()))
		return
	}

	var audit *model.Audit

	err1 := json.Unmarshal([]byte(result), audit)
	if err1 != nil {
		clog.Error("Worker: Failed to unmarshal audit log", clog.String("worker", worker.name))
		return
	}

	worker.app.Srv().Store.Audit().Save(audit)

	clog.Info("Worker: Job is complete", clog.String("worker", worker.name))
	worker.SetJobSuccess(job)
}

func (worker *Worker) SetJobSuccess(job *model.Job) {

}

func (woker *Worker) SetJobError(job *model.Job, appError *model.AppError) {

}
