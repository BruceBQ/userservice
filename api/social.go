package api

import (
	"encoding/json"
	"net/http"
	"userservice/model"
)

func (api *API) InitSocical() {
	api.BaseRoutes.Social.Handle("/login", api.ApiHandler(socialLogin)).Methods("POST")
	api.BaseRoutes.Social.Handle("/logout", api.ApiHandler(socialLogout)).Methods("POST")
	// api.BaseRoutes.Social.Handle("/@me", api.ApiHandler(getSocialMe)).Methods("GET")
	api.BaseRoutes.Social.Handle("/status", api.ApiHandler(updateSocialStatus)).Methods("PUT")
	api.BaseRoutes.Social.Handle("/{user_id}/plate_numbers", api.ApiHandler(getPlateNumbers)).Methods("GET")
	api.BaseRoutes.Social.Handle("/{user_id}/plate_numbers", api.ApiHandler(addPlateNumber)).Methods("POST")
	api.BaseRoutes.Social.Handle("/{user_id}/plate_numbers", api.ApiHandler(deletePlateNumber)).Methods("DELETE")
}

func socialLogin(c *Context, w http.ResponseWriter, r *http.Request) {
	user := model.SocialUserFromJson(r.Body)
	errmap := user.Validate()

	if len(errmap) > 0 {
		c.SetInvalidParams(errmap)
		return
	}

	if user == nil {
		return
	}

	ruser, err := c.App.SocialLogin(user)
	if err != nil {
		c.Err = err
		return
	}

	w.Write([]byte(ruser.ToJson()))

}

func socialLogout(c *Context, w http.ResponseWriter, r *http.Request) {
	if c.App.Session().ID != "" {
		if err := c.App.RevokeSessionById(c.App.Session().ID); err != nil {
			c.Err = err
			return
		}
	}

	w.Write([]byte("OK"))
}

func getSocialMe(c *Context, w http.ResponseWriter, r *http.Request) {
	c.RequireUserIdFromHeader(r.Header.Get("user_id"))

}

func updateSocialStatus(c *Context, w http.ResponseWriter, r *http.Request) {

}

func getPlateNumbers(c *Context, w http.ResponseWriter, r *http.Request) {
	c.RequireUserId()

	if c.Params.UserId != r.Header.Get("user_id") {
		c.Err = model.NewAppError("", "", "Unauthorized", nil, "", http.StatusUnauthorized)
		return
	}

	plateNumbers, err := c.App.GetPlateNumbers(c.Params.UserId)
	if err != nil {
		c.Err = err
		return
	}

	plates := struct {
		PlateNumber []*string `json:"plateNumbers"`
	}{
		PlateNumber: plateNumbers,
	}

	b, _ := json.Marshal(plates)
	w.Write([]byte(b))
}

func addPlateNumber(c *Context, w http.ResponseWriter, r *http.Request) {

	c.RequireUserId()
	if c.Params.UserId != r.Header.Get("user_id") {
		c.Err = model.NewAppError("", "", "Unauthorized", nil, "", http.StatusUnauthorized)
		return
	}

	props := model.MapFromJson(r.Body)
	plateNumber := props["plateNumber"]

	errmap := make(map[string]interface{})

	if len(plateNumber) == 0 {
		errmap["plateNumber"] = "Nhập biển số."
	}

	if len(errmap) > 0 {
		c.SetInvalidParams(errmap)
		return
	}

	err := c.App.AddPlateNumber(c.Params.UserId, plateNumber)

	if err != nil {
		c.Err = err
		return
	}

	w.Write([]byte("OK"))
}

func deletePlateNumber(c *Context, w http.ResponseWriter, r *http.Request) {
	c.RequireUserId()
	if c.Params.UserId != r.Header.Get("user_id") {
		c.Err = model.NewAppError("", "", "Unauthorized", nil, "", http.StatusUnauthorized)
		return
	}

	props := model.MapFromJson(r.Body)
	plateNumber := props["plateNumber"]

	errmap := make(map[string]interface{})

	if len(plateNumber) == 0 {
		errmap["plateNumber"] = "Nhập biển số."
	}

	if len(errmap) > 0 {
		c.SetInvalidParams(errmap)
		return
	}

	err := c.App.DeletePlateNumber(c.Params.UserId, plateNumber)
	if err != nil {
		c.Err = err
		return
	}

	w.Write([]byte("OK"))
}
