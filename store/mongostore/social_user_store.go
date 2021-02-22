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
)

type MongoSocialUserStore struct {
	MongoStore
	Client         *mongo.Client
	Db             *mongo.Database
	CollectionName string
}

func newMongoSocialUserStore(ms *MongoSupplier) store.SocialUserStore {
	s := &MongoSocialUserStore{
		Client:         ms.Client,
		Db:             ms.Db,
		CollectionName: "public_users",
	}
	return s
}

func (su MongoSocialUserStore) Create(user *model.SocialUser) (string, error) {
	ruser, err := su.Db.Collection(su.CollectionName).InsertOne(context.TODO(), user)
	if err != nil {
		return "", err
	}
	id := ruser.InsertedID.(primitive.ObjectID).Hex()
	return id, nil
}

func (su MongoSocialUserStore) GetById(id string) (*model.SocialUser, error) {
	var user model.SocialUser

	hexID, errID := primitive.ObjectIDFromHex(id)
	if errID != nil {
		return nil, errID
	}

	filter := bson.M{"_id": hexID}

	if err := su.Db.Collection(su.CollectionName).FindOne(context.Background(), filter).Decode(&user); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, store.NewErrNotFound("Social User", fmt.Sprintf("id=%s", id))
		}

		return nil, errors.Wrapf(err, "failed to get User with id=%s", id)
	}

	return &user, nil
}

func (su MongoSocialUserStore) GetByEmail(email string) (*model.SocialUser, error) {
	filter := bson.M{"email": email}
	user := model.SocialUser{}

	if err := su.Db.Collection(su.CollectionName).FindOne(context.Background(), filter).Decode(&user); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, store.NewErrNotFound("Social User", fmt.Sprintf("email=%s", email))
		}

		return nil, errors.Wrapf(err, "failed to get User with email=%s", email)
	}

	return &user, nil
}

func (su MongoSocialUserStore) AddPlateNumber(id string, plateNumber string) error {
	hexID, errID := primitive.ObjectIDFromHex(id)
	if errID != nil {
		return errID
	}

	filter := bson.M{"_id": hexID}

	_, err := su.Db.Collection(su.CollectionName).UpdateOne(context.Background(), filter, bson.M{
		"$addToSet": bson.M{
			"plateNumbers": plateNumber,
		},
	})

	return err
}

func (su MongoSocialUserStore) GetPlateNumbers(id string) ([]*string, error) {
	hexID, errID := primitive.ObjectIDFromHex(id)
	if errID != nil {
		return nil, errID
	}
	filter := bson.M{"_id": hexID}
	var user model.SocialUser

	err := su.Db.Collection(su.CollectionName).FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		return nil, err
	}

	return user.PlateNumbers, nil
}

func (su MongoSocialUserStore) DeletePlateNumber(id string, plateNumber string) error {
	hexID, errID := primitive.ObjectIDFromHex(id)
	if errID != nil {
		return errID
	}

	filter := bson.M{"_id": hexID}

	_, err := su.Db.Collection(su.CollectionName).UpdateOne(context.Background(), filter, bson.M{
		"$pull": bson.M{
			"plateNumbers": plateNumber,
		},
	})

	return err
}
