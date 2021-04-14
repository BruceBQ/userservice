package app

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"net/http"
	"time"
	"userservice/clog"
	"userservice/model"
	"userservice/store"
)

func (a *App) CreateSession(session *model.Session) (*model.Session, *model.AppError) {
	session.Token = ""

	session, err := a.Srv().Store.Session().Create(session)
	if err != nil {
		return nil, model.NewAppError("CreateSession", "app.session.create.app_error", "", nil, err.Error(), http.StatusInternalServerError)
	}

	a.AddSessionToCache(session)

	return session, nil
}

func (a *App) SetSessionExpireInDays(session *model.Session, days int) {
	if session.CreatedAt == 0 {
		session.ExpiresAt = model.GetMillis() + (1000 * 60 * 60 * 24 * int64(days))
	} else {
		session.ExpiresAt = session.CreatedAt + (1000 * 60 * 60 * 24 * int64(days))
	}
}

func (a *App) AddSessionToCache(session *model.Session) error {
	b, _ := json.Marshal(session)

	_, err := a.Srv().SessionCache.Set(context.Background(), "session:"+session.Token, b, 10*time.Minute).Result()
	if err != nil {
		fmt.Printf("Add to cache failed %s \n", err.Error())
	}

	return err
}

func (a *App) RevokeSession(session *model.Session) *model.AppError {
	if err := a.Srv().Store.Session().Delete(session.ID); err != nil {
		return model.NewAppError("RevokeSession", "app.session.delete.app_error", "Revoke session fail", nil, err.Error(), http.StatusInternalServerError)
	}

	a.ClearSessionCacheForUser(session.Token)
	return nil
}

func (a *App) RevokeSessionById(sessionId string) *model.AppError {
	session, err := a.Srv().Store.Session().Get(sessionId)
	if err != nil {
		return model.NewAppError("RevokeSessionById", "app.session.get.app_error", "", nil, err.Error(), http.StatusBadRequest)
	}
	return a.RevokeSession(session)
}

func (a *App) ClearSessionCacheForUser(token string) {
	_, err := a.Srv().SessionCache.Del(context.Background(), "session:"+token).Result()
	if err != nil {
		clog.Error("Failed to delete session in cache :" + err.Error())
	}

}

func (a *App) UpdateLastActivityAtIfNeeded(session model.Session) {
	now := model.GetMillis()

	if now-session.LastActivityAt < model.SESSION_ACTIVITY_TIMEOUT {
		return
	}

	err := a.Srv().Store.Session().UpdateLastActivityAt(session.ID, now)
	if err != nil {
		clog.Error("Failed to update LastActivityAt")
	}

	session.LastActivityAt = now
	a.AddSessionToCache(&session)
}

func (a *App) GetSession(token string) (*model.Session, *model.AppError) {
	// var session model.Session = &model.Session{}
	var session *model.Session

	result, err := a.Srv().SessionCache.Get(context.Background(), "session:"+token).Result()
	if err == nil {
		if err = json.Unmarshal([]byte(result), &session); err != nil {
			return nil, model.NewAppError("GetSession", "api.context.invalid_token.error", err.Error(), nil, "", http.StatusInternalServerError)
		}

		return session, nil
	}

	var nErr error
	session, nErr = a.Srv().Store.Session().Get(token)

	if nErr != nil {
		var nfErr *store.ErrNotFound
		switch {
		case errors.As(nErr, &nfErr):
			return nil, model.NewAppError("GetSession", "app.session.not_found", "Không tìm thấy session", nil, err.Error(), http.StatusBadRequest)
		default:
			return nil, model.NewAppError("GetSession", "app.session.app_error", "Lấy thông tin session thất bại.", nil, err.Error(), http.StatusInternalServerError)
		}
	}

	if session.IsExpired() {
		return nil, model.NewAppError("GetSession", "app.context.expired_token", "Phiên hết hạn", nil, "session is expired", http.StatusUnauthorized)
	}

	a.AddSessionToCache(session)
	return session, nil
}

func (a *App) ExtendSessionExpiryIfNeeded(session *model.Session) bool {
	if session == nil || session.IsExpired() {
		return false
	}

	sessionLength := a.GetSessionLengthInMillis(session)
	// Only extend the expiry if the lessor of 1% or 1 day has elapsed within the
	// current session duration.
	threshold := int64(math.Min(float64(sessionLength)*0.01, float64(24*60*60*1000)))

	if threshold < 5*60*1000 {
		threshold = 5 * 60 * 1000
	}

	now := model.GetMillis()
	elapsed := now - (session.ExpiresAt - sessionLength)
	if elapsed < threshold {
		return false
	}

	newExpiry := now + sessionLength

	if err := a.Srv().Store.Session().UpdateExpiresAt(session.ID, newExpiry); err != nil {
		return false
	}

	session.ExpiresAt = newExpiry
	a.AddSessionToCache(session)

	return true
}

func (a *App) GetSessionLengthInMillis(session *model.Session) int64 {
	if session == nil {
		return 0
	}

	days := *a.Config().ServiceSettings.SessionLengthWebInDays

	return int64(days * 24 * 60 * 60 * 1000)
}
