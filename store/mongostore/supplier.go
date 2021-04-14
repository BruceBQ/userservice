package mongostore

import (
	"context"
	"fmt"
	"os"
	"userservice/model"
	"userservice/store"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoSupplierStores struct {
	user       store.UserStore
	socialUser store.SocialUserStore
	page       store.PageStore
	permission store.PermissionStore
	role       store.RoleStore
	session    store.SessionStore
	audit      store.AuditStore
}

type MongoSupplier struct {
	Client   *mongo.Client
	Db       *mongo.Database
	settings *model.MongoSettings
	stores   MongoSupplierStores
	context  context.Context
}

func NewMongoSupplier(settings model.MongoSettings) *MongoSupplier {
	supplier := &MongoSupplier{
		settings: &settings,
	}

	supplier.initConnection()

	supplier.stores.user = newMongoUserStore(supplier)
	supplier.stores.socialUser = newMongoSocialUserStore(supplier)
	supplier.stores.page = newMongoPageStore(supplier)
	supplier.stores.permission = newMongoPermissionStore(supplier)
	supplier.stores.role = newMongoRoleStore(supplier)
	supplier.stores.session = newMongoSessionStore(supplier)
	supplier.stores.audit = newMongoAuditStore(supplier)

	return supplier
}

func (ms *MongoSupplier) initConnection() {
	client := setupConnection(ms.settings)

	ms.Client = client
	ms.Db = client.Database(*ms.settings.Database)
}

func setupConnection(setting *model.MongoSettings) *mongo.Client {
	uri := "mongodb://"
	if setting.Username != nil {
		uri += *setting.Username + ":" + *setting.Password + "@" + *setting.Address
	} else {
		uri += *setting.Address
	}

	clientOptions := options.Client().ApplyURI(uri)
	clientOptions.SetDirect(true)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		fmt.Printf("Connect to MongoDB error: %s\n", err.Error())
		os.Exit(1)
	}

	return client
}

func (ms *MongoSupplier) User() store.UserStore {
	return ms.stores.user
}

func (ms *MongoSupplier) SocialUser() store.SocialUserStore {
	return ms.stores.socialUser
}

func (ms *MongoSupplier) Page() store.PageStore {
	return ms.stores.page
}

func (ms *MongoSupplier) Permission() store.PermissionStore {
	return ms.stores.permission
}

func (ms *MongoSupplier) Role() store.RoleStore {
	return ms.stores.role
}

func (ms *MongoSupplier) Session() store.SessionStore {
	return ms.stores.session
}

func (ms *MongoSupplier) Audit() store.AuditStore {
	return ms.stores.audit
}
