package jobs

import (
	"errors"
	"time"
	"userservice/clog"
	"userservice/model"
)

type Schedulers struct {
	stop    chan bool
	stopped chan bool
	jobs    *JobServer
	running bool

	schedulers   []model.Scheduler
	nextRunTimes []*time.Time
}

var (
	ErrSchedulersNotRunning    = errors.New("job schedulers are not running")
	ErrSchedulersRunning       = errors.New("job schedulers are running")
	ErrSchedulersUninitialized = errors.New("job schedulers are not initialized")
)

func (srv *JobServer) InitSchedulders() error {
	srv.mut.Lock()
	defer srv.mut.Unlock()
	if srv.schedulers != nil && srv.schedulers.running {
		return ErrSchedulersRunning
	}
	clog.Debug("Initialising schedulers.")

	schedulers := &Schedulers{
		stop:    make(chan bool),
		stopped: make(chan bool),
		jobs:    srv,
	}

	if srv.AuditJob != nil {
		schedulers.schedulers = append(schedulers.schedulers, srv.AuditJob.MakeScheduler())
	}
	schedulers.nextRunTimes = make([]*time.Time, len(schedulers.schedulers))
	srv.schedulers = schedulers

	return nil
}

func (schedulers *Schedulers) Start() {
	go func() {
		clog.Info("Starting schedulers.")

		defer func() {
			clog.Info("Schedulers stopped.")
			close(schedulers.stopped)
		}()

		now := time.Now()
		for idx := range schedulers.schedulers {
			schedulers.SetNextRunTime(idx, now)
		}

		for {
			timer := time.NewTimer(10 * time.Second)
			select {
			case <-schedulers.stop:
				clog.Debug("Schedulers received stop signal.")
				timer.Stop()
				return
			case now = <-timer.C:
				clog.Debug("Scheduler Timeout")

				for idx, nextTime := range schedulers.nextRunTimes {
					scheduler := schedulers.schedulers[idx]

					job, err := schedulers.ScheduleJob(scheduler)
					if err != nil {
						continue
					}

					schedulers.jobs.workers.Audit.JobChannel() <- *job
					if nextTime == nil {
						continue
					}

					// if time.Now().After(*nextTime) {
					// 	scheduler := schedulers.schedulers[idx]
					// 	if scheduler == nil {
					// 		continue
					// 	}
					// 	schedulers.ScheduleJob(scheduler)
					// 	schedulers.SetNextRunTime(idx, now)
					// }

				}
			}

			timer.Stop()
		}
	}()

	schedulers.running = true
}

func (schedulers *Schedulers) Stop() {
	clog.Info("Stopping schedulers.")
	close(schedulers.stop)
	<-schedulers.stopped
	schedulers.running = false
}

func (schedulers *Schedulers) SetNextRunTime(idx int, now time.Time) {
	scheduler := schedulers.schedulers[idx]
	schedulers.nextRunTimes[idx] = scheduler.NextScheduleTime(now)
}

func (schedulers *Schedulers) ScheduleJob(scheduler model.Scheduler) (*model.Job, *model.AppError) {
	return scheduler.ScheduleJob()
}
