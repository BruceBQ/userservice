package web

import (
	"net/http"

	"github.com/gorilla/mux"
)

const (
	PAGE_DEFAULT     = 0
	PER_PAGE_DEFAULT = 50
)

type Params struct {
	UserId string
	RoleId string
}

func ParamsFromRequest(r *http.Request) *Params {
	params := &Params{}
	props := mux.Vars(r)
	// query := r.URL.Query()

	if val, ok := props["user_id"]; ok {
		params.UserId = val
	}

	if val, ok := props["role_id"]; ok {
		params.RoleId = val
	}

	return params
}
