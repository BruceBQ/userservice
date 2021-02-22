package app

import "userservice/model"

func (s *Server) Config() *model.Config {
	return s.configStore.Get()
}

func (a *App) Config() *model.Config {
	return a.Srv().Config()
}

func (a *App) ReloadConfig() error {
	return nil
}
