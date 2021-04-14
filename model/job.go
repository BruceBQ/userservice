package model

import "time"

const (
	JOB_TYPE_AUDIT = "audit"

	JOB_STATUS_PENDING          = "pending"
	JOB_STATUS_IN_PROGRESS      = "in_progress"
	JOB_STATUS_SUCCESS          = "success"
	JOB_STATUS_ERROR            = "error"
	JOB_STATUS_CANCEL_REQUESTED = "cancel_requested"
	JOB_STATUS_CANCELED         = "canceled"
	JOB_STATUS_WARNING          = "warning"
)

type Job struct {
	Id        string            `json:"id"`
	Type      string            `json:"type"`
	Data      map[string]string `json:"data"`
	CreatedAt int64             `json:"createdAt"`
}

type Worker interface {
	Run()
	Stop()
	JobChannel() chan<- Job
}

type Scheduler interface {
	Name() string
	NextScheduleTime(now time.Time) *time.Time
	ScheduleJob() (*Job, *AppError)
}
