package jobs

import (
	"sync"
	tjobs "userservice/jobs/interfaces"
	"userservice/store"
)

type JobServer struct {
	store store.Store

	AuditJob   tjobs.AuditInterface
	mut        sync.Mutex
	workers    *Workers
	schedulers *Schedulers
}

func NewJobServer(store store.Store) *JobServer {
	return &JobServer{
		store: store,
	}
}

func (srv *JobServer) StartWorkers() error {
	srv.mut.Lock()
	defer srv.mut.Unlock()

	srv.workers.Start()
	return nil
}

func (srv *JobServer) StartSchedulers() error {
	srv.mut.Lock()
	defer srv.mut.Unlock()
	if srv.schedulers == nil {
		return ErrSchedulersUninitialized
	} else if srv.schedulers.running {
		return ErrSchedulersRunning
	}

	srv.schedulers.Start()

	return nil
}
