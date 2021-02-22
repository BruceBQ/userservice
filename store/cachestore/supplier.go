package cachestore

import (
	"userservice/model"

	"github.com/go-redis/redis/v8"
)

type RedisSupplier struct {
	settings *model.RedisSettings
	Client   *redis.Client
}

func NewRedisClient(settings model.RedisSettings) *RedisSupplier {
	supplier := &RedisSupplier{
		settings: &settings,
	}

	supplier.init()
	return supplier
}

func (rs *RedisSupplier) init() {
	client := setupClient(rs.settings)
	rs.Client = client
}

func setupClient(settings *model.RedisSettings) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: *settings.Address,
		// Password: *settings.Password,
		DB: *settings.DB,
	})
	return client
}
