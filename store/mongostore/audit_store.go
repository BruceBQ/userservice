package mongostore

import (
	"context"
	"userservice/model"
	"userservice/store"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoAuditStore struct {
	MongoStore
	Client         *mongo.Client
	Db             *mongo.Database
	CollectionName string
}

func newMongoAuditStore(ms *MongoSupplier) store.AuditStore {
	return &MongoAuditStore{
		Client:         ms.Client,
		Db:             ms.Db,
		CollectionName: "audit_logs",
	}
}

func (as MongoAuditStore) Get(before int64, page int, userId string, permission string) ([]*model.Audit, error) {
	var audits []*model.Audit

	addFields := bson.D{{
		Key: "$addFields", Value: bson.D{
			{Key: "user_object_id", Value: bson.D{{Key: "$toObjectId", Value: "$user_id"}}},
		},
	}}

	condition := bson.D{{Key: "created_at", Value: bson.D{{Key: "$lte", Value: before}}}}
	userCondition := bson.D{{Key: "user_id", Value: userId}}
	permissionCondition := bson.D{{Key: "permission_name", Value: permission}}
	if len(userId) > 0 {
		condition = append(condition, userCondition...)
	}

	if len(permission) > 0 {
		condition = append(condition, permissionCondition...)
	}

	match := bson.D{{Key: "$match", Value: condition}}

	lookupUser := bson.D{{
		Key: "$lookup", Value: bson.D{
			{Key: "from", Value: "users"},
			{Key: "let", Value: bson.D{
				{Key: "user_object_id", Value: "$user_object_id"},
			}},
			{Key: "pipeline", Value: mongo.Pipeline{
				bson.D{
					{Key: "$match", Value: bson.D{
						{Key: "$expr", Value: bson.D{
							{Key: "$and", Value: []interface{}{
								bson.D{{Key: "$eq", Value: []interface{}{"$_id", "$$user_object_id"}}},
							}},
						}},
					}},
				},
				bson.D{
					{Key: "$project", Value: bson.D{
						{Key: "name", Value: 1},
					}},
				},
			}},
			{Key: "as", Value: "user"},
		},
	}}

	lookupPermission := bson.D{{
		Key: "$lookup", Value: bson.D{
			{Key: "from", Value: "permissions"},
			{Key: "let", Value: bson.D{
				{Key: "permission_name", Value: "$permission_name"},
			}},
			{Key: "pipeline", Value: mongo.Pipeline{
				bson.D{
					{Key: "$match", Value: bson.D{
						{Key: "$expr", Value: bson.D{
							{Key: "$and", Value: []interface{}{
								bson.D{{Key: "$eq", Value: []interface{}{"$name", "$$permission_name"}}},
							}},
						}},
					}},
				},
				bson.D{
					{Key: "$project", Value: bson.D{
						{Key: "_id", Value: 0},
						{Key: "name", Value: 1},
						{Key: "displayName", Value: 1},
					}},
				},
			}},
			{Key: "as", Value: "permission"},
		},
	}}

	project := bson.D{{
		Key: "$project", Value: bson.D{
			{Key: "user", Value: bson.D{{Key: "$arrayElemAt", Value: []interface{}{"$user", 0}}}},
			{Key: "permission", Value: bson.D{{Key: "$arrayElemAt", Value: []interface{}{"$permission", 0}}}},
			{Key: "data", Value: 1},
			{Key: "created_at", Value: 1},
		},
	}}

	sort := bson.D{{Key: "$sort", Value: bson.D{{Key: "created_at", Value: -1}}}}

	skip := bson.D{{Key: "$skip", Value: (page - 1) * model.AuditPerPage}}

	limit := bson.D{{Key: "$limit", Value: model.AuditPerPage}}

	cursor, err := as.Db.Collection(as.CollectionName).Aggregate(context.Background(), mongo.Pipeline{addFields, match, lookupUser, lookupPermission, project, sort, skip, limit})
	if err != nil {
		return nil, err
	}

	for cursor.Next(context.Background()) {
		var audit model.Audit
		cursor.Decode(&audit)
		audits = append(audits, &audit)
	}

	if audits == nil {
		audits = make([]*model.Audit, 0)
	}

	return audits, nil
}

func (as MongoAuditStore) Save(audit *model.Audit) error {
	audit.PreSave()
	_, err := as.Db.Collection(as.CollectionName).InsertOne(context.Background(), audit)
	if err != nil {
		return err
	}

	return nil
}

func (as MongoAuditStore) SaveMany(audits model.Audits) error {
	_, err := as.Db.Collection(as.CollectionName).InsertMany(context.Background(), []interface{}{})
	if err != nil {
		return err
	}

	return nil
}

func (as MongoAuditStore) Count(before int64, userId string, permission string) (int64, error) {
	condition := bson.D{{Key: "created_at", Value: bson.D{{Key: "$lte", Value: before}}}}
	userCondition := bson.D{{Key: "user_id", Value: userId}}
	permissionCondition := bson.D{{Key: "permission_name", Value: permission}}
	if len(userId) > 0 {
		condition = append(condition, userCondition...)
	}

	if len(permission) > 0 {
		condition = append(condition, permissionCondition...)
	}

	match := bson.D{{Key: "$match", Value: condition}}

	count := bson.D{{Key: "$count", Value: "count"}}

	cursor, err := as.Db.Collection(as.CollectionName).Aggregate(context.Background(), mongo.Pipeline{match, count})
	if err != nil {
		return -1, err
	}

	var result model.Count

	for cursor.Next(context.Background()) {
		cursor.Decode(&result)
	}

	return result.Count, nil
}
