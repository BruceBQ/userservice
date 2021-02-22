package api

import (
	"encoding/json"
	"net/http"
	"userservice/model"
)

func (api *API) InitRole() {
	api.BaseRoutes.Role.Handle("", api.ApiHandler(getAllRoles)).Methods("GET")
	api.BaseRoutes.Role.Handle("", api.ApiHandler(createRole)).Methods("POST")
	api.BaseRoutes.Role.Handle("/{role_id}", api.ApiHandler(getRole)).Methods("GET")
	api.BaseRoutes.Role.Handle("/{role_id}", api.ApiHandler(updateRole)).Methods("PUT")
	api.BaseRoutes.Role.Handle("/{role_id}", api.ApiHandler(deleteRole)).Methods("DELETE")
}

func getAllRoles(c *Context, w http.ResponseWriter, r *http.Request) {
	// if !c.App.UserHasPermissionTo(r.Header.Get("user_id"), model.PERMISSION_GET_ROLES) {
	// 	c.SetPermissionError(model.PERMISSION_GET_ROLES)
	// 	return
	// }

	roles, pages, err := c.App.GetRoles()

	if err != nil {

	}

	result := struct {
		Roles []*model.Role `json:"roles"`
		Pages []*model.Page `json:"pages"`
	}{
		Roles: roles,
		Pages: pages,
	}

	str, _ := json.Marshal(result)
	w.Write([]byte(string(str)))
}

func createRole(c *Context, w http.ResponseWriter, r *http.Request) {
	// if !c.App.UserHasPermissionTo(r.Header.Get("user_id"), model.PERMISSION_CREATE_ROLE) {
	// 	c.SetPermissionError(model.PERMISSION_CREATE_ROLE)
	// 	return
	// }

	role := model.RoleFromJSON(r.Body)

	errmap := role.Validate()

	if len(errmap) > 0 {
		c.SetInvalidParams(errmap)
		return
	}

	rrole, err := c.App.CreateRole(role)

	if err != nil {
		c.Err = err
		return
	}

	w.Write([]byte(rrole.ToJSON()))
}

func updateRole(c *Context, w http.ResponseWriter, r *http.Request) {
	// if !c.App.UserHasPermissionTo(r.Header.Get("user_id"), model.PERMISSION_UPDATE_ROLE) {
	// 	c.SetPermissionError(model.PERMISSION_UPDATE_ROLE)
	// 	return
	// }

	c.RequireRoleId()
	if c.Err != nil {
		return
	}

	role := model.RoleFromJSON(r.Body)

	errmap := role.Validate()
	if len(errmap) > 0 {
		c.SetInvalidParams(errmap)
		return
	}

	rrole, err := c.App.UpdateRole(c.Params.RoleId, role)

	if err != nil {
		c.Err = err
	}

	w.Write([]byte(rrole.ToJSON()))
}

func getRole(c *Context, w http.ResponseWriter, r *http.Request) {
	// if !c.App.UserHasPermissionTo(r.Header.Get("user_id"), model.PERMISSION_GET_ROLE) {
	// 	c.SetPermissionError(model.PERMISSION_GET_ROLE)
	// 	return
	// }

	c.RequireRoleId()
	if c.Err != nil {
		return
	}
	role, pages, permissions, err := c.App.GetRole(c.Params.RoleId)

	if err != nil {
		c.Err = err
		return
	}

	result := struct {
		Data        *model.Role         `json:"role"`
		Pages       []*model.Page       `json:"pages"`
		Permissions []*model.Permission `json:"permissions"`
	}{
		Data:        role,
		Pages:       pages,
		Permissions: permissions,
	}

	str, _ := json.Marshal(result)
	w.Write([]byte(string(str)))
}

func deleteRole(c *Context, w http.ResponseWriter, r *http.Request) {
	// if !c.App.UserHasPermissionTo(r.Header.Get("user_id"), model.PERMISSION_DELETE_ROLE) {
	// 	c.SetPermissionError(model.PERMISSION_DELETE_ROLE)
	// 	return
	// }

	c.RequireRoleId()
	if c.Err != nil {
		return
	}

	err := c.App.DeleteRole(c.Params.RoleId)

	if err != nil {
		c.Err = err
		return
	}

	w.Write([]byte("OK"))

}
