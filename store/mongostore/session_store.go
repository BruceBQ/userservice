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

type MongoSessionStore struct {
	MongoStore
	Client         *mongo.Client
	Db             *mongo.Database
	CollectionName string
}

func newMongoSessionStore(ms *MongoSupplier) store.SessionStore {
	ss := &MongoSessionStore{
		Client:         ms.Client,
		Db:             ms.Db,
		CollectionName: "sessions",
	}
	return ss
}

func (ss MongoSessionStore) Get(sessionIDOrToken string) (*model.Session, error) {
	filter := bson.M{}

	hexID, errID := primitive.ObjectIDFromHex(sessionIDOrToken)
	if errID == nil {
		filter = bson.M{"_id": hexID}
	} else {
		filter = bson.M{"token": sessionIDOrToken}
	}

	var session model.Session

	err := ss.Db.Collection(ss.CollectionName).FindOne(context.Background(), filter).Decode(&session)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, store.NewErrNotFound("Session", fmt.Sprintf("sessionIDOrToken=%s", sessionIDOrToken))
		}

		return nil, errors.Wrapf(err, "failed to get Session with sessionIDOrToken=%s", sessionIDOrToken)
	}

	return &session, nil

}

func (ss MongoSessionStore) Create(session *model.Session) (*model.Session, error) {
	session.PreSave()

	result, err := ss.Db.Collection(ss.CollectionName).InsertOne(context.Background(), session)
	if err != nil {
		return nil, err
	}

	session.ID = result.InsertedID.(primitive.ObjectID).Hex()
	return session, nil
}

func (ss MongoSessionStore) UpdateLastActivityAt(sessionId string, time int64) error {
	hexID, errID := primitive.ObjectIDFromHex(sessionId)
	if errID != nil {
		return errID
	}

	filter := bson.M{"_id": hexID}

	_, err := ss.Db.Collection(ss.CollectionName).UpdateOne(context.Background(), filter, bson.M{
		"$set": bson.M{
			"last_activity_at": time,
		},
	})

	return err
}

func (ss MongoSessionStore) UpdateExpiresAt(sessionID string, time int64) error {
	hexID, errID := primitive.ObjectIDFromHex(sessionID)
	if errID != nil {
		return errID
	}

	filter := bson.M{"_id": hexID}

	_, err := ss.Db.Collection(ss.CollectionName).UpdateOne(context.Background(), filter, bson.M{
		"$set": bson.M{
			"expires_at": time,
		},
	})

	return err
}

func (ss MongoSessionStore) Delete(sessionIDOrToken string) error {
	filter := bson.M{}

	hexID, errID := primitive.ObjectIDFromHex(sessionIDOrToken)
	if errID == nil {
		filter = bson.M{"_id": hexID}
	} else {
		filter = bson.M{"token": sessionIDOrToken}
	}

	_, err := ss.Db.Collection(ss.CollectionName).DeleteOne(context.Background(), filter)
	return err
}
