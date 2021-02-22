package cache

import (
	"errors"
	"time"
)

var ErrKeyNotFound = errors.New("key not found")

type Cache interface {
	Purge() error

	Set(key string, value interface{}) error

	SetWithDefaultExpiry(key string, value interface{}) error

	SetWithExpiry(key string, value interface{}, ttl time.Duration) error

	Get(key string, value interface{}) error

	Remove(key string) error

	Keys() ([]string, error)

	Len() (int, error)

	Name() string
}
