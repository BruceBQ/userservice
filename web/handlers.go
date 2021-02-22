package web

import (
	"net/http"
	"userservice/app"
	"userservice/clog"
	"userservice/model"
)

type Handler struct {
	GetGlobalAppOptions app.AppOptionCreator
	HandleFunc          func(*Context, http.ResponseWriter, *http.Request)
	RequireSession      bool
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w = newWrappedWriter(w)

	var statusCode int
	defer func() {
		responseFields := []clog.Field{
			clog.String("method", r.Method),
			clog.String("url", r.URL.Path),
			clog.Int("status_code", statusCode),
		}
		clog.Debug("Received HTTP request", responseFields...)
	}()

	c := &Context{}

	c.App = app.New(h.GetGlobalAppOptions()...)
	c.App.InitServer()

	c.Params = ParamsFromRequest(r)
	c.Log = c.App.Log()

	w.Header().Set("Content-Type", "application/json")

	token := app.ParseAuthTokenFromRequest(r)
	if token != "" {
		session, err := c.App.GetSession(token)

		if err != nil {
			if err.StatusCode == http.StatusInternalServerError {
				c.Err = err
			} else if h.RequireSession {
				c.Err = model.NewAppError("ServeHTTP", "api.context.session_expired.app_error", "Session is expired or illegal", nil, "token="+token, http.StatusUnauthorized)
			}
		} else {
			c.App.SetSession(session)
		}
	}

	if c.Err == nil {
		h.HandleFunc(c, w, r)
	}

	// Handle errors that have occurred
	if c.Err != nil {
		c.Err.Where = r.URL.Path
		w.WriteHeader(c.Err.StatusCode)
		c.Err.Sanitize(map[string]bool{})
		w.Write([]byte(c.Err.ToJson()))
	}

	statusCode = w.(*responseWriterWrapper).StatusCode()
}
