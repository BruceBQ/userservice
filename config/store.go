package config

import (
	"userservice/model"
)

type Store interface {
	Get() *model.Config
	Set(*model.Config) (*model.Config, error)
}

func NewStore(dsn string) (Store, error) {
	return NewFileStore(dsn)
}
