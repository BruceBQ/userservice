package app

import (
	"context"
	"userservice/clog"
	"userservice/model"
)

type App struct {
	srv *Server

	session model.Session
	context context.Context
}

func New(options ...AppOption) *App {
	app := &App{}

	for _, option := range options {
		option(app)
	}

	return app
}

func (a *App) InitServer() {
	a.srv.AppInitializedOnce.Do(func() {
		// a.InitJobs()
		// a.Srv().RunJobs()
	})
}

func (a *App) Srv() *Server {
	return a.srv
}

func (a *App) Log() *clog.Logger {
	return a.srv.Log
}

func (a *App) Session() *model.Session {
	return &a.session
}

func (a *App) SetSession(s *model.Session) {
	a.session = *s
}

func (a *App) InitJobs() {
	if jobsAuditInterface != nil {
		a.srv.Jobs.AuditJob = jobsAuditInterface(a)
	}

	// a.srv.Jobs.InitWorkers()
	// a.srv.Jobs.InitSchedulders()
}
