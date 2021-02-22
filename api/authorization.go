package api

import (
	"net/http"
	"userservice/model"
)

func (api *API) InitAuthorization() {
	api.BaseRoutes.Auth.Handle("", api.ApiHandler(authorize)).Methods("POST")
}

func authorize(c *Context, w http.ResponseWriter, r *http.Request) {
	props := model.MapFromJson(r.Body)

	token := props["token"]

	if token == "" {
		c.Err = model.NewAppError("api.authorize", "", "Xác thực thất bại.", nil, "", http.StatusUnauthorized)
		return
	}

	session, err := c.App.GetSession(token)

	if err != nil {
		c.Err = err
		return
	}

	w.Write([]byte(session.ToJSON()))

}
