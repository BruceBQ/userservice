package app

import (
	"net/http"
	"userservice/model"
)

func (a *App) MakeAuthenticationError() *model.AppError {
	return model.NewAppError("Authentication", "api.context.authentication.app_error", "Unauthorized", nil, "", http.StatusUnauthorized)
}

func (a *App) checkUserPassword(user *model.User, password string) *model.AppError {
	if !model.ComparePassword(user.Password, password) {
		return model.NewAppError("checkUserPassword", "api.user.check_user_password.invalid.app_error", "Password does not match", map[string]interface{}{"password": "Mật khẩu không khớp!"}, "", http.StatusUnauthorized)
	}

	return nil
}

func ParseAuthTokenFromRequest(r *http.Request) string {
	authHeader := r.Header.Get(model.HEADER_AUTH)

	return authHeader
}
