package app

import (
	"errors"
	"fmt"
	"net/http"
	"userservice/model"
	"userservice/store"
)

func (a *App) GetAllUsers() ([]*model.User, *model.AppError) {
	users, err := a.Srv().Store.User().GetAll()
	if err != nil {
		return nil, model.NewAppError("GetAllUsers", "app.user.get_all_users", "Lấy danh sách người dùng thất bại.", nil, err.Error(), http.StatusInternalServerError)
	}

	return users, nil
}

func (a *App) GetUser(userId string) (*model.User, *model.AppError) {
	user, err := a.Srv().Store.User().Get(userId)
	if err != nil {
		var nfError *store.ErrNotFound
		switch {
		case errors.As(err, &nfError):
			return nil, model.NewAppError("GetUser", "app.user.get_user.app_error", "Không tìm thấy thông tin người dùng.", nil, err.Error(), http.StatusBadRequest)
		default:
			return nil, model.NewAppError("GetUser", "app.user.get_user.internal_error", "Lấy thông tin người dùng thất bại.", nil, err.Error(), http.StatusInternalServerError)
		}
	}

	return user, nil
}

func (a *App) CreateUser(user *model.User) (*model.User, *model.AppError) {
	// role, _, _, errRole := a.GetRole(user.RoleID)
	err := a.checkCamera(user.Cameras)
	if err != nil {
		return nil, err
	}

	_, errRole := a.Srv().Store.Role().GetByID(user.RoleID)
	if errRole != nil {
		return nil, model.NewAppError("CreateUser", "app.user.create_user.app_error", "Nhóm người dùng không hợp lệ.", map[string]interface{}{"role": "Nhóm người dùng không hợp lệ."}, errRole.Error(), http.StatusBadRequest)
	}

	// if role.BuiltIn && role.Name == model.SUPER_ADMIN_ROLE {
	// 	return nil, model.NewAppError("CreateUser", "app.user.create_user.app_error", "Nhóm người dùng không hợp lệ.", map[string]interface{}{"role": "Nhóm người dùng không hợp lệ."}, "", http.StatusBadRequest)
	// }

	_, errEmail := a.Srv().Store.User().GetByEmail(user.Email)
	if errEmail != nil {
		var nfError *store.ErrNotFound
		switch {
		case errors.As(errEmail, &nfError):
			id, err := a.Srv().Store.User().Create(user)
			if err != nil {
				return nil, model.NewAppError("CreateUser", "app.user.create_user.internal_error", "Tạo người dùng thất bại", nil, err.Error(), http.StatusInternalServerError)
			}
			ruser, err := a.Srv().Store.User().Get(id)
			if err != nil {
				return nil, model.NewAppError("CreateUser", "app.user.create_user.internal_error", "Không thể lấy thông tin người dùng sau khi tạo", nil, err.Error(), http.StatusInternalServerError)
			}

			return ruser, nil
		default:
			return nil, model.NewAppError("CreateUser", "app.user.create_user.app_error", "Không thể tạo người dùng", nil, "", http.StatusInternalServerError)
		}
	}

	return nil, model.NewAppError("CreateUser", "app.user.create_user.app_error", "Email đã tồn tại.", map[string]interface{}{"email": "Email đã tồn tại"}, "", http.StatusBadRequest)
}

func (a *App) UpdateUser(id string, user *model.User) (*model.User, *model.AppError) {
	err := a.checkCamera(user.Cameras)
	if err != nil {
		return nil, err
	}

	_, errRole := a.Srv().Store.Role().GetByID(user.RoleID)
	if errRole != nil {
		return nil, model.NewAppError("UpdateUser", "app.user.update_user.app_error", "Nhóm người dùng không hợp lệ.", map[string]interface{}{"role": "Nhóm người dùng không hợp lệ."}, errRole.Error(), http.StatusBadRequest)
	}
	// if role.BuiltIn && role.Name == model.SUPER_ADMIN_ROLE {
	// 	return nil, model.NewAppError("UpdateUser", "app.user.update_user.app_error", "Nhóm người dùng không hợp lệ.", map[string]interface{}{"role": "Nhóm người dùng không hợp lệ."}, "", http.StatusBadRequest)
	// }

	errUpdate := a.Srv().Store.User().Update(id, user)
	if errUpdate != nil {
		var nfError *store.ErrNotFound
		switch {
		case errors.As(errUpdate, &nfError):
			return nil, model.NewAppError("UpdateUser", "app.user.update_user.app_error", "Không tìm thấý người dùng", nil, errUpdate.Error(), http.StatusBadRequest)
		default:
			return nil, model.NewAppError("UpdateUser", "app.user.update_user.internal_error", "Cập nhật thông tin người dùng thất bại.", nil, errUpdate.Error(), http.StatusInternalServerError)
		}
	}

	ruser, errUser := a.Srv().Store.User().Get(id)
	if errUser != nil {
		return nil, model.NewAppError("UpdateRole", "app.user.update_user.internal_error", "Không thể lấy thông tin người dùng sau khi cập nhật.", nil, errUser.Error(), http.StatusInternalServerError)
	}

	return ruser, nil
}

func (a *App) DeleteUser(id string) *model.AppError {
	user, err := a.Srv().Store.User().Get(id)
	if err != nil {
		return model.NewAppError("DeleteUser", "app.user.delete_user.app_error", "Không thể lấy thông tin người dùng này.", nil, err.Error(), http.StatusInternalServerError)
	}

	if user.Role.Name == model.SUPER_ADMIN_ROLE {
		return model.NewAppError("DeleteUser", "app.user.delete_user.app_error", "Không thể xóa người dùng này.", nil, "User is super admin", http.StatusBadRequest)
	}

	err = a.Srv().Store.User().Delete(id)
	if err != nil {
		var nfError *store.ErrNotFound
		switch {
		case errors.As(err, &nfError):
			return model.NewAppError("DeleteUser", "app.user.delete_user.app_error", "Không tìm thấy người dùng.", nil, err.Error(), http.StatusBadRequest)
		default:
			return model.NewAppError("DeleteUser", "app.user.delete_user.internal_error", "Xóa người dùng thất bại.", nil, err.Error(), http.StatusInternalServerError)
		}
	}

	return nil
}

func (a *App) UpdateUserInfo(id string, userInfo *model.UserInfo) (*model.User, *model.AppError) {
	user, err := a.Srv().Store.User().UpdateInfo(id, userInfo)

	if err != nil {

	}

	return user, nil
}

func (a *App) UpdateUserPassword(userID string, password string, newPassword string) *model.AppError {
	user, err := a.Srv().Store.User().Get(userID)
	fmt.Printf("User %v \n", user)
	if err != nil {
		return model.NewAppError("UpdateUserPassword", "app.user.get.app_error", "Không thể lấy thông tin người dùng", nil, err.Error(), http.StatusInternalServerError)
	}

	if err1 := a.checkUserPassword(user, password); err1 != nil {
		return err1
	}

	err2 := a.Srv().Store.User().UpdatePassword(userID, newPassword)
	if err2 != nil {
		return model.NewAppError("UpdateUserPassword", "app.user.update.app_error", "Cập nhật mật khẩu thất bại", nil, err.Error(), http.StatusInternalServerError)
	}

	return nil
}

func (a *App) GetMe(userID string) (*model.User, *model.AppError) {
	user, err := a.Srv().Store.User().GetRolePermission(userID)
	if err != nil {
		return nil, nil
	}

	user.Sanitize(map[string]bool{})

	return user, nil
}
