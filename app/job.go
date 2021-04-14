package app

import (
	tjobs "userservice/jobs/interfaces"
)

var jobsAuditInterface func(*App) tjobs.AuditInterface

func RegisterJobsAuditInterface(f func(*App) tjobs.AuditInterface) {
	jobsAuditInterface = f
}
