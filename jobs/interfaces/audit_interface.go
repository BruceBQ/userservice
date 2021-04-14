package interfaces

import "userservice/model"

type AuditInterface interface {
	MakeWorker() model.Worker
	MakeScheduler() model.Scheduler
}
