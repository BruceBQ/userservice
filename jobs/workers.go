package jobs

import (
	"errors"
	"userservice/clog"
	"userservice/model"
)

type Workers struct {
	Audit   model.Worker
	running bool
}

var (
	ErrWorkersNotRunning   = errors.New("jobs workers are not running")
	ErrWorkersRunning      = errors.New("Job worker are running")
	ErrWorkerUninitialized = errors.New("Job workers are not initialized")
)

func (srv *JobServer) InitWorkers() error {
	srv.mut.Lock()
	defer srv.mut.Unlock()

	if srv.workers != nil {
		return ErrWorkersRunning
	}

	workers := &Workers{}

	if srv.AuditJob != nil {
		workers.Audit = srv.AuditJob.MakeWorker()
	}

	srv.workers = workers
	return nil
}

func (workers *Workers) Start() {
	clog.Info("Starting workers")
	if workers.Audit != nil {
		go workers.Audit.Run()
	}

	workers.running = true
}

func (workers *Workers) Stop() {
	if workers.Audit != nil {
		workers.Audit.Stop()
	}

	workers.running = false

	clog.Info("Stopped workers")
}
