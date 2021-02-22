package app

import (
	"userservice/config"
)

type Option func(s *Server) error

func ConfigStore(configStore config.Store) Option {
	return func(s *Server) error {
		s.configStore = configStore
		return nil
	}
}

type AppOption func(a *App)
type AppOptionCreator func() []AppOption

func ServerConnector(s *Server) AppOption {
	return func(a *App) {
		a.srv = s
	}
}
