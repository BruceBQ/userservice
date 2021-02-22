package config

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	defaultHost     = "127.0.0.1:27017"
	defaultDatabase = "camera"
)

type DatabaseStore struct {
	Client *mongo.Client
	Db     *mongo.Database
}

func NewDatabaseStore(uri string) (ds *DatabaseStore, err error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	ds = &DatabaseStore{
		Client: client,
		Db:     client.Database("camera"),
	}

	return ds, nil
}
