package mongostore

import "userservice/store"

type MongoStore interface {
	User() store.UserStore
	SocialUser() store.SocialUserStore
	Page() store.PageStore
	Permission() store.PermissionStore
	Role() store.RoleStore
}
