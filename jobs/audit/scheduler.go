package audit

import (
	"time"
	"userservice/app"
	"userservice/model"
)

const (
	ScheFreqSeconds = 3
)

type Scheduler struct {
	App *app.App
}

func (m *AuditJobInterfaceImpl) MakeScheduler() model.Scheduler {
	return &Scheduler{m.App}
}

func (scheduler *Scheduler) Name() string {
	return JobName + "Scheduler"
}

func (scheduler *Scheduler) NextScheduleTime(now time.Time) *time.Time {
	nextTime := time.Now().Add(ScheFreqSeconds * time.Second)
	return &nextTime
}

func (scheduler *Scheduler) ScheduleJob() (*model.Job, *model.AppError) {
	data := map[string]string{}
	job, err := scheduler.App.Srv().Jobs.CreateJob(model.JOB_TYPE_AUDIT, data)
	if err != nil {
		return nil, err
	}

	return job, nil
}
