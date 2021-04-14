package store

import (
	"userservice/model"
)

type Store interface {
	User() UserStore
	SocialUser() SocialUserStore
	Page() PageStore
	Permission() PermissionStore
	Role() RoleStore
	Session() SessionStore
	Audit() AuditStore
}

type UserStore interface {
	GetAll() ([]*model.User, error)
	Get(id string) (*model.User, error)
	GetByUsername(username string) (*model.User, *model.AppError)
	Create(*model.User) (string, error)
	GetByEmail(string) (*model.User, error)
	Update(string, *model.User) error
	UpdateInfo(string, *model.UserInfo) (*model.User, error)
	UpdatePassword(string, string) error
	Delete(string) error
	GetRolePermission(string) (*model.User, error)
	DeleteCameraFromUser(string) error
	GetFilterAudit() ([]*model.User, error)
}

type SocialUserStore interface {
	Create(*model.SocialUser) (string, error)
	GetById(string) (*model.SocialUser, error)
	GetByEmail(string) (*model.SocialUser, error)
	AddPlateNumber(string, string) error
	GetPlateNumbers(string) ([]*string, error)
	DeletePlateNumber(string, string) error
}

type PageStore interface {
	Get() ([]*model.Page, error)
}

type PermissionStore interface {
	GetAll() ([]*model.Permission, error)
	GetAdmin() ([]*model.Permission, error)
	GetPublic() ([]*model.Permission, error)
	GetByName(string) (*model.Permission, error)
	GetFilterAudit() ([]*model.Permission, error)
}

type RoleStore interface {
	GetAll() ([]*model.Role, error)
	GetByID(string) (*model.Role, error)
	GetByName(string) (*model.Role, error)
	Create(*model.Role) (string, error)
	Update(string, *model.Role) error
	Delete(string) error
}

type SessionStore interface {
	Get(sessionIDOrToken string) (*model.Session, error)
	Create(*model.Session) (*model.Session, error)
	UpdateLastActivityAt(string, int64) error
	UpdateExpiresAt(string, int64) error
	Delete(string) error
}

type AuditStore interface {
	Get(before int64, page int, userId string, permission string) ([]*model.Audit, error)
	Save(*model.Audit) error
	SaveMany(model.Audits) error
	Count(int64, string, string) (int64, error)
}
