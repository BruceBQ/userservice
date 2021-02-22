package mongostore

import (
	"context"
	"fmt"
	"userservice/model"
	"userservice/store"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoPermissionStore struct {
	MongoStore
	Client         *mongo.Client
	Db             *mongo.Database
	CollectionName string
}

func newMongoPermissionStore(ms *MongoSupplier) store.PermissionStore {
	ps := &MongoPermissionStore{
		Client:         ms.Client,
		Db:             ms.Db,
		CollectionName: "permissions",
	}

	return ps
}

func (ps MongoPermissionStore) GetAll() ([]*model.Permission, error) {
	projection := bson.D{{Key: "_id", Value: 0}}
	filter := bson.M{}

	cursor, err := ps.Db.Collection(ps.CollectionName).Find(context.Background(), filter, options.Find().SetProjection(projection))
	if err != nil {
		return nil, err
	}

	var permissions []*model.Permission

	for cursor.Next(context.TODO()) {
		var permission model.Permission

		err := cursor.Decode(&permission)
		if err != nil {
			return nil, err
		}

		permissions = append(permissions, &permission)
	}

	return permissions, nil
}

func (ps MongoPermissionStore) GetAdmin() ([]*model.Permission, error) {
	projection := bson.D{{Key: "_id", Value: 0}, {Key: "name", Value: 1}, {Key: "displayName", Value: 1}, {Key: "path", Value: 1}, {Key: "description", Value: 1}, {Key: "method", Value: 1}, {Key: "type", Value: 1}}
	filter := bson.M{"type": "admin"}

	cursor, err := ps.Db.Collection(ps.CollectionName).Find(context.Background(), filter, options.Find().SetProjection(projection))
	if err != nil {
		return nil, err
	}

	var permissions []*model.Permission

	for cursor.Next(context.TODO()) {
		var permission model.Permission

		err := cursor.Decode(&permission)
		if err != nil {
			return nil, err
		}
		// permission.Sanitize()
		permissions = append(permissions, &permission)
	}

	return permissions, nil
}

func (ps MongoPermissionStore) GetPublic() ([]*model.Permission, error) {
	projection := bson.D{{Key: "_id", Value: 0}, {Key: "name", Value: 1}, {Key: "displayName", Value: 1}, {Key: "path", Value: 1}, {Key: "description", Value: 1}, {Key: "method", Value: 1}, {Key: "type", Value: 1}}
	filter := bson.M{"type": "public"}

	cursor, err := ps.Db.Collection(ps.CollectionName).Find(context.Background(), filter, options.Find().SetProjection(projection))
	if err != nil {
		return nil, err
	}

	var permissions []*model.Permission

	for cursor.Next(context.TODO()) {
		var permission model.Permission

		err := cursor.Decode(&permission)
		if err != nil {
			return nil, err
		}
		// permission.Sanitize()
		permissions = append(permissions, &permission)
	}

	return permissions, nil
}

func (ps MongoPermissionStore) GetByName(name string) (*model.Permission, error) {
	filter := bson.M{"name": name}

	var permission model.Permission

	err := ps.Db.Collection(ps.CollectionName).FindOne(context.Background(), filter).Decode(&permission)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, store.NewErrNotFound("Permission", fmt.Sprintf("name=%s", name))
		}

		return nil, errors.Wrapf(err, "failed to get Permission with name=%s", name)
	}

	return &permission, nil
}
