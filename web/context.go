package web

import (
	"fmt"
	"net/http"
	"userservice/app"
	"userservice/clog"
	"userservice/model"
)

type Context struct {
	App    *app.App
	Log    *clog.Logger
	Params *Params
	Err    *model.AppError
}

func (c *Context) SetInvalidParams(params map[string]interface{}) {
	c.Err = NewInvalidParamsError(params)
}

func (c *Context) SetUnauthorized() {
	// c.Err = model.NewAppError("", "", "Un")
}

func (c *Context) RequireUserId() *Context {
	if c.Err != nil {
		return c
	}

	if !model.IsMongoId(c.Params.UserId) {
		c.SetInvalidParams(map[string]interface{}{"userId": "Invalid"})
	}

	return c
}

func (c *Context) RequireUserIdFromHeader(userID string) *Context {
	fmt.Printf("Check user id header %s \n", userID)
	if c.Err != nil {
		return c
	}

	if !model.IsMongoId(userID) {
		c.SetInvalidParams(map[string]interface{}{"userId": "Invalid"})
		fmt.Printf("Error 1 %v \n", c.Err)
	}
	fmt.Printf("Error 2 %v \n", c.Err)
	return c
}

func (c *Context) RequireRoleId() *Context {
	if c.Err != nil {
		return c
	}

	if !model.IsMongoId(c.Params.RoleId) {
		c.SetInvalidParams(map[string]interface{}{"roleId": "Invalid"})
	}

	return c
}
func (c *Context) SetPermissionError(permission *model.Permission) {
	c.Err = c.App.MakePermissionError(permission)
}

func NewInvalidParamsError(params map[string]interface{}) *model.AppError {
	err := model.NewAppError("Context", "api.context.invalid_body_param", "Invalid form body", params, "", http.StatusBadRequest)
	return err
}
