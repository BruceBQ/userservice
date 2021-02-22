package mongostore

import (
	"context"
	"userservice/model"
	"userservice/store"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoPageStore struct {
	MongoStore
	Client         *mongo.Client
	Db             *mongo.Database
	CollectionName string
}

func newMongoPageStore(ms *MongoSupplier) store.PageStore {
	ps := &MongoPageStore{
		Client:         ms.Client,
		Db:             ms.Db,
		CollectionName: "pages",
	}

	return ps
}

func (ps MongoPageStore) Get() ([]*model.Page, error) {
	cursor, err := ps.Db.Collection(ps.CollectionName).Find(context.Background(), bson.M{})

	if err != nil {
		return nil, err
	}

	defer cursor.Close(context.Background())

	var pages []*model.Page

	for cursor.Next(context.TODO()) {
		var page model.Page

		err := cursor.Decode(&page)
		if err != nil {
			return nil, err
		}

		pages = append(pages, &page)
	}

	return pages, nil
}
