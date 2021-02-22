package store

type BaseStore struct {
	Store
	UserStore       UserStore
	SocialUserStore SocialUserStore
	PageStore       PageStore
	PermissionStore PermissionStore
	RoleStore       RoleStore
	SessionStore    SessionStore
}

func (s *BaseStore) User() UserStore {
	return s.UserStore
}

func (s *BaseStore) SocialUser() SocialUserStore {
	return s.SocialUserStore
}

func (s *BaseStore) Page() PageStore {
	return s.PageStore
}

func (s *BaseStore) Permission() PermissionStore {
	return s.PermissionStore
}

func (s *BaseStore) Role() RoleStore {
	return s.RoleStore
}

func (s *BaseStore) Session() SessionStore {
	return s.SessionStore
}

func New(childStore Store) *BaseStore {
	newStore := BaseStore{}

	newStore.UserStore = childStore.User()
	newStore.SocialUserStore = childStore.SocialUser()
	newStore.PageStore = childStore.Page()
	newStore.PermissionStore = childStore.Permission()
	newStore.RoleStore = childStore.Role()
	newStore.SessionStore = childStore.Session()
	return &newStore
}
