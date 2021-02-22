package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"userservice/model"
)

func (api *API) InitUser() {
	api.BaseRoutes.Users.Handle("/login", api.ApiHandler(login)).Methods("POST")
	api.BaseRoutes.Users.Handle("/logout", api.ApiHandler(logout)).Methods("POST")
	api.BaseRoutes.Users.Handle("", api.ApiHandler(createUser)).Methods("POST")
	api.BaseRoutes.Users.Handle("", api.ApiHandler(getAllUsers)).Methods("GET")
	api.BaseRoutes.Users.Handle("/status", api.ApiSessionRequired(updateStatus)).Methods("PUT")
	api.BaseRoutes.Users.Handle("/@me", api.ApiHandler(getMe)).Methods("GET")
	api.BaseRoutes.Users.Handle("/{user_id}", api.ApiHandler(getUser)).Methods("GET")
	api.BaseRoutes.Users.Handle("/{user_id}", api.ApiHandler(updateUser)).Methods("PUT")
	api.BaseRoutes.Users.Handle("/{user_id}/info", api.ApiHandler(updateInfo)).Methods("PUT")
	api.BaseRoutes.Users.Handle("/{user_id}/password", api.ApiHandler((updatePassword))).Methods("PUT")
	api.BaseRoutes.Users.Handle("/{user_id}", api.ApiHandler(deleteUser)).Methods("DELETE")
}

func createUser(c *Context, w http.ResponseWriter, r *http.Request) {
	user := model.UserFromJSON(r.Body)

	errmap := user.Validate()

	if len(errmap) > 0 {
		c.SetInvalidParams(errmap)
		return
	}

	ruser, err := c.App.CreateUser(user)
	if err != nil {
		c.Err = err
		return
	}

	ruser.Sanitize(make(map[string]bool))

	w.Write([]byte(string(ruser.ToJson())))
}

func getAllUsers(c *Context, w http.ResponseWriter, r *http.Request) {
	// if !c.App.UserHasPermissionTo(r.Header.Get("user_id"), model.PERMISSION_GET_USERS) {
	// 	c.SetPermissionError(model.PERMISSION_GET_USERS)
	// 	return
	// }

	users, err := c.App.GetAllUsers()

	if err != nil {
		c.Err = err
		return
	}

	str, _ := json.Marshal(users)
	w.Write([]byte(str))
}

func login(c *Context, w http.ResponseWriter, r *http.Request) {
	props := model.MapFromJson(r.Body)
	email := props["email"]
	password := props["password"]

	errmap := make(map[string]interface{})

	if !model.IsValidEmail(email) {
		errmap["email"] = "Email không hợp lệ!"
	}

	if len(email) == 0 {
		errmap["email"] = "Nhập email!"
	}

	if len(password) == 0 {
		errmap["password"] = "Nhập mật khẩu!"
	}

	if len(errmap) != 0 {
		c.SetInvalidParams(errmap)
		return
	}

	user, err := c.App.DoLogin(email, password)

	if err != nil {
		c.Err = err
		return
	}

	// user.Sanitize(map[string]bool{})

	w.Write([]byte(user.ToJSON()))
}

func logout(c *Context, w http.ResponseWriter, r *http.Request) {
	if c.App.Session().ID != "" {
		if err := c.App.RevokeSessionById(c.App.Session().ID); err != nil {
			c.Err = err
			return
		}
	}

	w.Write([]byte("OK"))
}

func getUser(c *Context, w http.ResponseWriter, r *http.Request) {
	// if !c.App.UserHasPermissionTo(r.Header.Get("user_id"), model.PERMISSION_GET_USER) {
	// 	c.SetPermissionError(model.PERMISSION_GET_USER)
	// 	return
	// }

	user, err := c.App.GetUser(c.Params.UserId)

	if err != nil {
		c.Err = err
		return
	}

	user.Sanitize(map[string]bool{})

	w.Write([]byte(user.ToJson()))
}

func getMe(c *Context, w http.ResponseWriter, r *http.Request) {
	// if r.Header.Get("user_id") == "" {
	// 	c.Err = model.NewAppError("", "", "empty user", nil, "", http.StatusBadGateway)
	// 	return
	// }

	c.RequireUserIdFromHeader(r.Header.Get("user_id"))

	ruser, err := c.App.GetMe(r.Header.Get("user_id"))
	if err != nil {
		c.Err = err
		return
	}

	w.Write([]byte(ruser.ToJson()))
}

func deleteUser(c *Context, w http.ResponseWriter, r *http.Request) {
	// if !c.App.UserHasPermissionTo(r.Header.Get("user_id"), model.PERMISSION_DELETE_ROLE) {
	// 	c.SetPermissionError(model.PERMISSION_DELETE_ROLE)
	// 	return
	// }

	c.RequireUserId()

	err := c.App.DeleteUser(c.Params.UserId)
	if err != nil {
		c.Err = err
		return
	}
}

func updateUser(c *Context, w http.ResponseWriter, r *http.Request) {
	// if !c.App.UserHasPermissionTo(r.Header.Get("user_id"), model.PERMISSION_UPDATE_ROLE) {
	// 	c.SetPermissionError(model.PERMISSION_UPDATE_ROLE)
	// 	return
	// }

	c.RequireUserId()

	user := model.UserFromJSON(r.Body)

	errmap := user.Validate()

	if len(errmap) > 0 {
		c.SetInvalidParams(errmap)
		return
	}

	ruser, err := c.App.UpdateUser(c.Params.UserId, user)

	if err != nil {
		c.Err = err
		return
	}

	w.Write([]byte(ruser.ToJson()))

}

func updateInfo(c *Context, w http.ResponseWriter, r *http.Request) {
	c.RequireUserId()

	if c.Params.UserId != r.Header.Get("user_id") {

	}

	userInfo := model.UserInfoFromJSON(r.Body)

	errmap := userInfo.Validate()
	if len(errmap) > 0 {
		c.SetInvalidParams(errmap)
		return
	}

	user, err := c.App.UpdateUserInfo(c.Params.UserId, userInfo)
	if err != nil {
		c.Err = err
		return
	}

	user.Sanitize(map[string]bool{})

	w.Write([]byte(user.ToJson()))
}

func updatePassword(c *Context, w http.ResponseWriter, r *http.Request) {
	c.RequireUserId()

	props := model.MapFromJson(r.Body)

	password := props["password"]
	newPassword := props["newPassword"]
	newPasswordConfirmation := props["newPasswordConfirmation"]

	errmap := make(map[string]interface{})

	if len(password) == 0 {
		errmap["password"] = "Nhập mật khẩu cũ."
	}

	if len(newPassword) == 0 {
		errmap["newPassword"] = "Nhập mật khẩu mới."
	}

	if len(newPassword) < model.USER_PASSWORD_MIN_LENGTH || len(newPassword) > model.USER_PASSWORD_MAX_LENGTH {
		errmap["newPassword"] = "Mật khẩu mới có tối thiểu " + strconv.Itoa(model.USER_PASSWORD_MIN_LENGTH) + " ký tự và tối đa " + strconv.Itoa(model.USER_PASSWORD_MAX_LENGTH) + " ký tự."
	}

	if newPasswordConfirmation != newPassword {
		errmap["newPasswordConfirmation"] = "Mật khẩu mới không khớp với xác nhận."
	}

	if len(newPasswordConfirmation) == 0 {
		errmap["newPasswordConfirmation"] = "Nhập xác nhận mật khẩu mới."
	}

	if len(errmap) > 0 {
		c.SetInvalidParams(errmap)
		return
	}

	err := c.App.UpdateUserPassword(c.Params.UserId, password, newPassword)
	if err != nil {
		c.Err = err
		return
	}

	w.Write([]byte("OK"))

}

func updateStatus(c *Context, w http.ResponseWriter, r *http.Request) {
	c.App.UpdateLastActivityAtIfNeeded(*c.App.Session())
	c.App.ExtendSessionExpiryIfNeeded(c.App.Session())

	w.Write([]byte("OK"))
}
