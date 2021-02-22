package mongostore

import (
	"context"
	"fmt"
	"userservice/model"
	"userservice/store"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoRoleStore struct {
	MongoStore
	Client         *mongo.Client
	Db             *mongo.Database
	CollectionName string
}

func newMongoRoleStore(ms *MongoSupplier) store.RoleStore {
	rs := &MongoRoleStore{
		Client:         ms.Client,
		Db:             ms.Db,
		CollectionName: "roles",
	}

	return rs
}

// GetAll Get all role
func (rs MongoRoleStore) GetAll() ([]*model.Role, error) {
	projection := bson.D{{Key: "name", Value: 1}, {Key: "description", Value: 1}, {Key: "pages", Value: 1}, {Key: "builtin", Value: 1}}

	cursor, err := rs.Db.Collection(rs.CollectionName).Find(context.Background(), bson.M{}, options.Find().SetProjection(projection))

	if err != nil {
		return nil, err
	}

	var roles []*model.Role

	for cursor.Next(context.Background()) {
		var role model.Role

		err := cursor.Decode(&role)
		if err != nil {
			return nil, err
		}

		roles = append(roles, &role)
	}

	return roles, nil
}

// GetByID Get role by id
func (rs MongoRoleStore) GetByID(id string) (*model.Role, error) {
	var role model.Role

	hexID, errID := primitive.ObjectIDFromHex(id)
	if errID != nil {
		return nil, errID
	}

	filter := bson.M{"_id": hexID}

	err := rs.Db.Collection(rs.CollectionName).FindOne(context.Background(), filter).Decode(&role)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, store.NewErrNotFound("Role", fmt.Sprintf("id=%s", id))
		}
		return nil, errors.Wrapf(err, "failed to get Role with id=%s", id)
	}

	return &role, nil
}

// GetByName Get role by name
func (rs MongoRoleStore) GetByName(name string) (*model.Role, error) {
	var role model.Role

	filter := bson.M{"name": name}
	err := rs.Db.Collection(rs.CollectionName).FindOne(context.Background(), filter).Decode(&role)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, store.NewErrNotFound("Role", fmt.Sprintf("name=%s", name))
		}

		return nil, errors.Wrapf(err, "failed to get Role with name=%s", name)
	}

	return &role, nil
}

// Create Create role
func (rs MongoRoleStore) Create(role *model.Role) (string, error) {
	role.PreSave()

	rrole, err := rs.Db.Collection(rs.CollectionName).InsertOne(context.Background(), role)
	if err != nil {
		return "", err
	}

	id := rrole.InsertedID.(primitive.ObjectID).Hex()

	return id, nil
}

// Update Update role
func (rs MongoRoleStore) Update(id string, role *model.Role) error {
	role.PreUpdate()

	hexID, errID := primitive.ObjectIDFromHex(id)
	if errID != nil {
		return errID
	}

	filter := bson.M{"_id": hexID}
	result, err := rs.Db.Collection(rs.CollectionName).UpdateOne(context.Background(), filter, bson.M{
		"$set": bson.M{
			"name":        role.Name,
			"description": role.Description,
			"pages":       role.Pages,
			"permissions": role.Permissions,
			"updated_at":  role.UpdatedAt,
		},
	})

	if err != nil {
		return errors.Wrapf(err, "failed to update Role with id=%s", id)
	}

	if result.ModifiedCount == 0 {
		return store.NewErrNotFound("Role", fmt.Sprintf("id=%s", id))
	}

	return nil
}

// Delete Delete role by id
func (rs MongoRoleStore) Delete(id string) error {
	hexID, errID := primitive.ObjectIDFromHex(id)
	if errID != nil {
		return errID
	}
	filter := bson.M{"_id": hexID, "builtin": bson.M{"$ne": true}}

	result, err := rs.Db.Collection(rs.CollectionName).DeleteOne(context.Background(), filter)

	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return errors.New("no role deleted")
	}

	return nil
}
