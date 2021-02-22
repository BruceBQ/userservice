package app

import (
	"net/http"
	"userservice/clog"
	"userservice/model"
)

func (a *App) MakePermissionError(permission *model.Permission) *model.AppError {
	return model.NewAppError("Permissions", "api.context.permissions.app_error", "Missing Access!", nil, "", http.StatusForbidden)
}

// func (a *App) SessionHasAdminPermission(session model.Session, path string, method string) (*model.User, *model.AppError) {
// 	user, err := a.Srv().Store.User().GetRolePermission(session.UserID)

// 	if err != nil {
// 		return nil, model.NewAppError("SessionHasAdminPermission", "app.authorization.get.app_error", "Get User failed", nil, err.Error(), http.StatusInternalServerError)
// 	}

// 	permissions, err := a.Srv().Store.Permission().GetAdmin()
// 	if err != nil {
// 		return nil, model.NewAppError("SessionHasAdminPermission", "app.authorization.get.app_error", "Get Permissions failed", nil, err.Error(), http.StatusInternalServerError)
// 	}

// 	newPath := strings.Replace(path, model.API_URL_SUFFIX, "", 1)
// 	if !a.RoleGrantPermission(user.Role, permissions, newPath, method) {
// 		return nil, model.NewAppError("SessionHasAdminPermission", "app.authorization.get.app_error", "Missing Access", nil, "", http.StatusForbidden)
// 	}

// 	return user, nil
// }

// func (a *App) SessionHasPublicPermission(session model.Session, path string, method string) (*model.SocialUser, *model.AppError) {
// 	user, err := a.Srv().Store.SocialUser().GetById(session.UserID)

// 	if err != nil {
// 		return nil, model.NewAppError("SessionHasPublicPermission", "app.authorization.get.app_error", "Get User failed", nil, err.Error(), http.StatusInternalServerError)
// 	}

// 	permissions, err := a.Srv().Store.Permission().GetPublic()
// 	if err != nil {
// 		return nil, model.NewAppError("SessionHasPublicPermission", "app.authorization.get.app_error", "Get Permissions failed", nil, "", http.StatusInternalServerError)
// 	}

// 	newPath := strings.Replace(path, model.API_URL_SUFFIX, "", 1)

// 	if !a.SocialUserGrantPermission(permissions, newPath, method) {
// 		return nil, model.NewAppError("SessionHasPublicPermission", "app.authorization.get.app_error", "Missing Access!", nil, "", http.StatusForbidden)
// 	}

// 	return user, nil

// }

// func (a *App) RoleGrantPermission(role *model.Role, permissions []*model.Permission, path string, method string) bool {
// 	for _, permission := range permissions {
// 		if permission.Method == method {
// 			matched, err := regexp.Match(permission.Path, []byte(path))
// 			if err != nil {
// 				return false
// 			}
// 			if matched {
// 				if _, ok := role.Permissions[permission.Name]; ok {
// 					return true
// 				}
// 			}
// 		}
// 	}

// 	return false
// }

// func (a *App) SocialUserGrantPermission(permissions []*model.Permission, path string, method string) bool {
// 	for _, permission := range permissions {
// 		if permission.Method == method {
// 			matched, err := regexp.Match(permission.Path, []byte(path))
// 			if err != nil {
// 				return false
// 			}

// 			if matched {
// 				return true
// 			}
// 		}
// 	}
// 	return false
// }

func (a *App) SessionHasPermissionTo(session model.Session, permission *model.Permission) bool {

	return false
}

func (a *App) UserHasPermissionTo(userID string, permission *model.Permission) bool {
	user, err := a.Srv().Store.User().GetRolePermission(userID)

	if err != nil {
		clog.Error("Get user for authorization failed userID " + userID + ", " + err.Error())
		return false
	}

	if a.RoleGrantPermission(user.Role, permission) {
		return true
	}

	return false
}

func (a *App) RoleGrantPermission(role *model.Role, permission *model.Permission) bool {
	if _, ok := role.Permissions[permission.Name]; ok {
		return true
	}

	return false
}

func (a *App) SessionHasPermissionToCamera(session model.Session, cameraID string) *model.AppError {
	return nil
}

func (a *App) HasPermissionToCamera(userId string, cameraId string) bool {
	return false
}
