package mongostore

import (
	"context"
	"fmt"
	"net/http"
	"userservice/clog"
	"userservice/model"
	"userservice/store"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoUserStore struct {
	MongoStore
	Client         *mongo.Client
	Db             *mongo.Database
	ctx            context.Context
	CollectionName string
}

func newMongoUserStore(ms *MongoSupplier) store.UserStore {
	us := &MongoUserStore{
		Client:         ms.Client,
		Db:             ms.Db,
		CollectionName: "users",
	}

	return us
}

func (us MongoUserStore) GetAll() ([]*model.User, error) {
	var users []*model.User

	addFields := bson.D{{
		Key: "$addFields", Value: bson.D{
			{Key: "role_object_id", Value: bson.D{{Key: "$toObjectId", Value: "$role_id"}}},
		},
	}}
	lookup := bson.D{{
		Key: "$lookup", Value: bson.D{
			{Key: "from", Value: "roles"},
			{Key: "let", Value: bson.D{{Key: "role_object_id", Value: "$role_object_id"}}},
			{Key: "pipeline", Value: mongo.Pipeline{
				bson.D{
					{Key: "$match", Value: bson.D{
						{Key: "$expr", Value: bson.D{
							{Key: "$and", Value: []interface{}{
								bson.D{{Key: "$eq", Value: []interface{}{"$_id", "$$role_object_id"}}},
							}},
						}},
					}},
				},
				bson.D{
					{Key: "$project", Value: bson.D{{Key: "name", Value: 1}, {Key: "description", Value: 1}}},
				},
			}},
			{Key: "as", Value: "role"},
		},
	}}
	project := bson.D{{
		Key: "$project", Value: bson.D{
			{Key: "email", Value: 1},
			{Key: "phone", Value: 1},
			{Key: "name", Value: 1},
			{Key: "description", Value: 1},
			{Key: "workplace", Value: 1},
			{Key: "role", Value: bson.D{{Key: "$arrayElemAt", Value: []interface{}{"$role", 0}}}},
			{Key: "created_at", Value: 1},
			{Key: "updated_at", Value: 1},
			{Key: "builtin", Value: 1},
		},
	}}

	cursor, err := us.Db.Collection(us.CollectionName).Aggregate(context.Background(), mongo.Pipeline{addFields, lookup, project})
	if err != nil {
		return nil, err
	}

	for cursor.Next(context.Background()) {
		var user model.User
		cursor.Decode(&user)
		users = append(users, &user)
	}

	return users, nil
}

func (us MongoUserStore) Get(id string) (*model.User, error) {
	var users []*model.User

	hexID, errID := primitive.ObjectIDFromHex(id)
	if errID != nil {
		return nil, errID
	}

	match := bson.D{{Key: "$match", Value: bson.D{{Key: "_id", Value: hexID}}}}
	addFields := bson.D{{
		Key: "$addFields", Value: bson.D{
			{Key: "role_object_id", Value: bson.D{{Key: "$toObjectId", Value: "$role_id"}}},
		},
	}}
	lookup := bson.D{{
		Key: "$lookup", Value: bson.D{
			{Key: "from", Value: "roles"},
			{Key: "let", Value: bson.D{{Key: "role_object_id", Value: "$role_object_id"}}},
			{Key: "pipeline", Value: mongo.Pipeline{
				bson.D{
					{Key: "$match", Value: bson.D{
						{Key: "$expr", Value: bson.D{
							{Key: "$and", Value: []interface{}{
								bson.D{{Key: "$eq", Value: []interface{}{"$_id", "$$role_object_id"}}},
							}},
						}},
					}},
				},
				bson.D{
					{Key: "$project", Value: bson.D{{Key: "name", Value: 1}, {Key: "description", Value: 1}}},
				},
			}},
			{Key: "as", Value: "role"},
		},
	}}
	project := bson.D{{
		Key: "$project", Value: bson.D{
			{Key: "email", Value: 1},
			{Key: "phone", Value: 1},
			{Key: "name", Value: 1},
			{Key: "password", Value: 1},
			{Key: "description", Value: 1},
			{Key: "workplace", Value: 1},
			{Key: "cameras", Value: 1},
			{Key: "role", Value: bson.D{{Key: "$arrayElemAt", Value: []interface{}{"$role", 0}}}},
			{Key: "last_password_update", Value: 1},
			{Key: "created_at", Value: 1},
			{Key: "updated_at", Value: 1},
			{Key: "builtin", Value: 1},
		},
	}}

	cursor, err := us.Db.Collection(us.CollectionName).Aggregate(context.Background(), mongo.Pipeline{addFields, match, lookup, project})

	if err != nil {
		return nil, err
	}

	for cursor.Next(context.Background()) {
		var user model.User

		err := cursor.Decode(&user)
		if err != nil {
			return nil, err
		}

		users = append(users, &user)
	}

	if len(users) == 0 {
		return nil, store.NewErrNotFound("Role", fmt.Sprintf("id=%s", id))
	}

	return users[0], nil
}

func (us MongoUserStore) GetByUsername(username string) (*model.User, *model.AppError) {
	var user model.User

	filter := bson.M{"username": username}
	err := us.Db.Collection("users").FindOne(context.TODO(), filter).Decode(&user)

	if err != nil {
		clog.Debug(err.Error())
		return nil, model.NewAppError("MongoUserStore.GetByUsername", "store.mongo_user.app_error", "", nil, err.Error(), http.StatusInternalServerError)
	}

	return &user, nil
}

func (us MongoUserStore) GetByEmail(email string) (*model.User, error) {
	var user model.User

	filter := bson.M{"email": email}
	err := us.Db.Collection(us.CollectionName).FindOne(context.Background(), filter).Decode(&user)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, store.NewErrNotFound("User", fmt.Sprintf("email=%s", email))
		}

		return nil, errors.Wrapf(err, "failed to get User with email=%s", email)
	}

	return &user, nil
}

func (us MongoUserStore) Create(user *model.User) (string, error) {
	user.PreSave()

	ruser, err := us.Db.Collection(us.CollectionName).InsertOne(context.Background(), user)

	if err != nil {
		return "", err
	}

	id := ruser.InsertedID.(primitive.ObjectID).Hex()
	return id, err

}

func (us MongoUserStore) Update(id string, user *model.User) error {
	user.PreUpdate()

	hexID, errID := primitive.ObjectIDFromHex(id)
	if errID != nil {
		return errID
	}

	filter := bson.M{"_id": hexID}

	result, err := us.Db.Collection(us.CollectionName).UpdateOne(context.Background(), filter, bson.M{
		"$set": bson.M{
			"name":        user.Name,
			"description": user.Description,
			"phone":       user.Phone,
			"workplace":   user.Workplace,
			"role_id":     user.RoleID,
			"cameras":     user.Cameras,
			"updated_at":  user.UpdatedAt,
		},
	})

	if err != nil {
		return errors.Wrapf(err, "failed to update User with id=%s", id)
	}

	if result.ModifiedCount == 0 {
		return store.NewErrNotFound("User", fmt.Sprintf("id=%s", id))
	}

	return nil
}

func (us MongoUserStore) UpdateInfo(id string, userInfo *model.UserInfo) (*model.User, error) {
	userInfo.PreUpdate()
	hexID, errID := primitive.ObjectIDFromHex(id)
	if errID != nil {
		return nil, errID
	}

	filter := bson.M{"_id": hexID}
	result, err := us.Db.Collection(us.CollectionName).UpdateOne(context.Background(), filter, bson.M{
		"$set": bson.M{
			"name":       userInfo.Name,
			"phone":      userInfo.Phone,
			"workplace":  userInfo.Workplace,
			"updated_at": userInfo.UpdatedAt,
		},
	})

	if err != nil {
		return nil, errors.Wrapf(err, "failed to update User Info with id=%s", id)
	}

	if result.ModifiedCount == 0 {
		return nil, store.NewErrNotFound("User", fmt.Sprintf("id=%s", id))
	}

	var user model.User

	err = us.Db.Collection(us.CollectionName).FindOne(context.Background(), filter).Decode(&user)

	if err != nil {
		return nil, errors.Wrapf(err, "failed to get User with id=%s", id)
	}

	return &user, nil
}

func (us MongoUserStore) UpdatePassword(id string, newPassword string) error {
	hexID, errID := primitive.ObjectIDFromHex(id)
	if errID != nil {
		return errID
	}

	filter := bson.M{"_id": hexID}

	result, err := us.Db.Collection(us.CollectionName).UpdateOne(context.Background(), filter, bson.M{
		"$set": bson.M{
			"password":             model.HashPassword(newPassword),
			"last_password_update": model.GetMillis(),
		},
	})

	if err != nil {
		errors.Wrapf(err, "failed to update User Password with id=%s", id)
	}

	if result.ModifiedCount == 0 {
		return store.NewErrNotFound("User", fmt.Sprintf("id=%s", id))
	}

	return nil
}

func (us MongoUserStore) Delete(id string) error {
	hexID, errID := primitive.ObjectIDFromHex(id)
	if errID != nil {
		return errID
	}

	filter := bson.M{"_id": hexID, "builtin": false}

	result, err := us.Db.Collection(us.CollectionName).DeleteOne(context.Background(), filter)

	if result.DeletedCount == 0 {
		return store.NewErrNotFound("User", fmt.Sprintf("id=%s", id))
	}

	return err
}

func (us MongoUserStore) GetRolePermission(id string) (*model.User, error) {
	hexID, errID := primitive.ObjectIDFromHex(id)
	if errID != nil {
		return nil, errID
	}

	match := bson.D{{Key: "$match", Value: bson.D{{Key: "_id", Value: hexID}}}}
	addFields := bson.D{{
		Key: "$addFields", Value: bson.D{
			{Key: "role_object_id", Value: bson.D{{Key: "$toObjectId", Value: "$role_id"}}},
		},
	}}
	lookup := bson.D{{
		Key: "$lookup", Value: bson.D{
			{Key: "from", Value: "roles"},
			{Key: "localField", Value: "role_object_id"},
			{Key: "foreignField", Value: "_id"},
			{Key: "as", Value: "role"},
		},
	}}
	project := bson.D{{
		Key: "$project", Value: bson.D{
			{Key: "email", Value: 1},
			{Key: "phone", Value: 1},
			{Key: "name", Value: 1},
			{Key: "password", Value: 1},
			{Key: "description", Value: 1},
			{Key: "workplace", Value: 1},
			{Key: "role", Value: bson.D{{Key: "$arrayElemAt", Value: []interface{}{"$role", 0}}}},
			{Key: "last_password_update", Value: 1},
			{Key: "created_at", Value: 1},
			{Key: "updated_at", Value: 1},
			{Key: "builtin", Value: 1},
		},
	}}

	var users []*model.User

	cursor, err := us.Db.Collection(us.CollectionName).Aggregate(context.Background(), mongo.Pipeline{match, addFields, lookup, project})
	if err != nil {
		return nil, err
	}

	for cursor.Next(context.Background()) {
		var user model.User

		err := cursor.Decode(&user)
		if err != nil {
			return nil, err
		}

		users = append(users, &user)
	}

	return users[0], nil
}

func (us MongoUserStore) DeleteCameraFromUser(cameraId string) error {
	_, err := us.Db.Collection(us.CollectionName).UpdateMany(context.Background(), bson.M{}, bson.M{
		"$pull": bson.M{
			"cameras": cameraId,
		},
	})

	return err
}
