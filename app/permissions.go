package app

import (
	"net/http"
	"userservice/model"
)

func (a *App) GetPermission() ([]*model.Page, []*model.Permission, *model.AppError) {
	pages, err := a.Srv().Store.Page().Get()
	if err != nil {
		return nil, nil, model.NewAppError("GetPermission", "", "Get page failed", nil, "", http.StatusInternalServerError)
	}
	permissions, err := a.Srv().Store.Permission().GetAll()
	if err != nil {
		return nil, nil, model.NewAppError("GetPermission", "", "Get permissions failed", nil, "", http.StatusInternalServerError)
	}

	return pages, permissions, nil
}
