package api

import (
	"encoding/json"
	"net/http"
	"userservice/model"
)

func (api *API) InitPermission() {
	api.BaseRoutes.Permission.Handle("", api.ApiHandler(getPermissions)).Methods("GET")
}

func getPermissions(c *Context, w http.ResponseWriter, r *http.Request) {

	pages, permissions, err := c.App.GetPermission()
	if err != nil {
		c.Err = err
		return
	}

	result := struct {
		Pages       []*model.Page       `json:"pages"`
		Permissions []*model.Permission `json:"permissions"`
	}{
		Pages:       pages,
		Permissions: permissions,
	}

	str, _ := json.Marshal(result)

	w.Write([]byte(string(str)))
}
