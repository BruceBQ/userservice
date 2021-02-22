package app

import (
	"userservice/clog"
	"userservice/model"
)

type AppIface interface {
	InitServer()
	Srv() *Server
	Log() *clog.Logger

	DoLogin(email string, password string) (*model.UserWithSession, *model.AppError)
	GetAllUsers() ([]*model.User, *model.AppError)
	CreateUser(*model.User) (*model.User, *model.AppError)
	GetUser(id string) (*model.User, *model.AppError)
	UpdateUser(string, *model.User) (*model.User, *model.AppError)
	DeleteUser(string) *model.AppError
	UpdateUserInfo(string, *model.UserInfo) (*model.User, *model.AppError)
	UpdateUserPassword(string, string, string) *model.AppError
	GetMe(string) (*model.User, *model.AppError)

	SocialLogin(*model.SocialUser) (*model.SocialUserWithSession, *model.AppError)
	AddPlateNumber(string, string) *model.AppError
	GetPlateNumbers(string) ([]*string, *model.AppError)
	DeletePlateNumber(string, string) *model.AppError

	MakePermissionError(*model.Permission) *model.AppError
	SessionHasAdminPermission(session model.Session, path string, method string) (*model.User, *model.AppError)
	SessionHasPublicPermission(session model.Session, path string, method string) (*model.SocialUser, *model.AppError)
	UserHasPermissionTo(string, *model.Permission) bool

	GetPermission() ([]*model.Page, []*model.Permission, *model.AppError)

	GetRoles() ([]*model.Role, []*model.Page, *model.AppError)
	CreateRole(*model.Role) (*model.Role, *model.AppError)
	GetRole(string) (*model.Role, []*model.Page, []*model.Permission, *model.AppError)
	UpdateRole(string, *model.Role) (*model.Role, *model.AppError)
	DeleteRole(string) *model.AppError

	Session() *model.Session
	SetSession(*model.Session)
	GetSession(string) (*model.Session, *model.AppError)
	UpdateLastActivityAtIfNeeded(model.Session)
	ExtendSessionExpiryIfNeeded(*model.Session) bool
	RevokeSessionById(string) *model.AppError
}
