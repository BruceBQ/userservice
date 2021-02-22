package app

import (
	"errors"
	"net/http"
	"userservice/model"
	"userservice/store"
)

func (a *App) DoLogin(email string, password string) (*model.UserWithSession, *model.AppError) {
	user, err := a.Srv().Store.User().GetByEmail(email)

	if err != nil {
		var nfError *store.ErrNotFound
		switch {
		case errors.As(err, &nfError):
			return nil, model.NewAppError("DoLogin", "app.login.dologin.app_error", "Đăng nhập không thành công", nil, err.Error(), http.StatusBadRequest)
		default:
			return nil, model.NewAppError("DoLogin", "app.login.dologin.internal_error", "Đăng nhập không thành công", nil, err.Error(), http.StatusInternalServerError)
		}
	}

	if err := a.checkUserPassword(user, password); err != nil {
		return nil, err
	}

	session := &model.Session{
		UserID: user.ID,
	}

	a.SetSessionExpireInDays(session, *a.Config().ServiceSettings.SessionLengthWebInDays)

	_, err1 := a.CreateSession(session)
	if err1 != nil {
		return nil, err1
	}

	user.Sanitize(map[string]bool{})

	ruser := &model.UserWithSession{
		User:  nil,
		Token: session.Token,
	}

	return ruser, nil
}
