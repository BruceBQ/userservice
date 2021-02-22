package api

import "net/http"

func (api *API) InitStatus() {
	api.BaseRoutes.Users.Handle("/{user_id}/status", api.ApiHandler(updateUserStatus)).Methods("PUT")
}

func updateUserStatus(c *Context, w http.ResponseWriter, r *http.Request) {
	c.RequireUserId()

	if c.Err != nil {
		return
	}
}
