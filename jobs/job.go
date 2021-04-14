package jobs

import "userservice/model"

func (srb *JobServer) CreateJob(jobType string, jobData map[string]string) (*model.Job, *model.AppError) {
	job := model.Job{
		Id:        "",
		Type:      jobType,
		CreatedAt: model.GetMillis(),
		Data:      jobData,
	}

	return &job, nil
}
