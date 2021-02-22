package app

import (
	"errors"
	"net/http"
	"userservice/model"
	"userservice/store"
)

func (a *App) GetRoles() ([]*model.Role, []*model.Page, *model.AppError) {
	roles, err := a.Srv().Store.Role().GetAll()
	if err != nil {

	}

	pages, err := a.Srv().Store.Page().Get()
	if err != nil {

	}

	return roles, pages, nil
}

func (a *App) CreateRole(role *model.Role) (*model.Role, *model.AppError) {
	rrole, err := a.Srv().Store.Role().GetByName(role.Name)

	if err != nil {
		var nfErr *store.ErrNotFound
		switch {
		case errors.As(err, &nfErr):
			id, err := a.Srv().Store.Role().Create(role)
			if err != nil {
				return nil, model.NewAppError("CreateRole", "app.role.create_role.internal_error", "Tạo nhóm người dùng thất bại", nil, "", http.StatusInternalServerError)
			}

			rrole, err = a.Srv().Store.Role().GetByID(id)

			if err != nil {
				return nil, model.NewAppError("CreateRole", "app.role.create_role.internal_error", "Tạo nhóm người dùng thất bại", nil, "", http.StatusInternalServerError)
			}

			return rrole, nil

		default:
			return nil, model.NewAppError("CreateRole", "app.role.create_role.app_error", "Tạo nhóm người dùng thất bại.", nil, err.Error(), http.StatusInternalServerError)
		}
	}

	return nil, model.NewAppError("CreateRole", "app.role.create_role.app_error", "Tạo nhóm người dùng thất bại.", map[string]interface{}{"name": "Tên nhóm người dùng đã tồn tại"}, "", http.StatusBadRequest)
}

func (a *App) GetRole(id string) (*model.Role, []*model.Page, []*model.Permission, *model.AppError) {
	role, err := a.Srv().Store.Role().GetByID(id)

	if err != nil {
		var nfErr *store.ErrNotFound
		switch {
		case errors.As(err, &nfErr):
			return nil, nil, nil, model.NewAppError("GetRole", "app.role.get_role.app_error", "Không tìm thấy nhóm người dùng.", nil, err.Error(), http.StatusBadRequest)
		default:
			return nil, nil, nil, model.NewAppError("GetRole", "app.role.get_role.internal_error", "Tải thông tin nhóm người dùng thất bại.", nil, "", http.StatusInternalServerError)
		}
	}

	pages, permissions, err1 := a.GetPermission()

	if err1 != nil {
		return nil, nil, nil, err1
	}

	return role, pages, permissions, nil
}

func (a *App) UpdateRole(id string, role *model.Role) (*model.Role, *model.AppError) {
	err := a.Srv().Store.Role().Update(id, role)

	if err != nil {
		return nil, model.NewAppError("Update Role", "app.role.update_role.internal_error", "Cập nhật thông tin nhóm người dùng thất bại.", nil, err.Error(), http.StatusInternalServerError)
	}

	rrole, err := a.Srv().Store.Role().GetByID(id)

	if err != nil {
		return nil, model.NewAppError("Update Role", "app.role.update_role.internal_error", "Cập nhật thông tin nhóm người dùng thất bại.", nil, err.Error(), http.StatusInternalServerError)
	}

	return rrole, nil
}

func (a *App) DeleteRole(id string) *model.AppError {
	role, err := a.Srv().Store.Role().GetByID(id)

	if err != nil {
		return model.NewAppError("DeleteRole", "app.role.delete_role.internal_error", "Xóa nhóm người dùng thất bại.", nil, "", http.StatusInternalServerError)
	}

	if role.BuiltIn {
		return model.NewAppError("DeleteRole", "app.role.delete_role.internal_error", "Không thể xóa nhóm người dùng này.", nil, "", http.StatusBadRequest)
	}

	err = a.Srv().Store.Role().Delete(id)

	if err != nil {
		return model.NewAppError("DeleteRole", "app.role.delete_role.internal_error", "Xóa nhóm người dùng thất bại.", nil, "", http.StatusInternalServerError)
	}

	return nil
}
