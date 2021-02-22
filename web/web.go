package web

import (
	"net/http"
	"userservice/app"
	"userservice/clog"

	"github.com/gorilla/mux"
)

type Web struct {
	GetGlobalAppOptions app.AppOptionCreator
	MainRouter          *mux.Router
}

func New(globalOptions app.AppOptionCreator, root *mux.Router) *Web {
	clog.Debug("Initializing web routes")

	web := &Web{
		GetGlobalAppOptions: globalOptions,
		MainRouter:          root,
	}

	return web
}

func Handle404(w http.ResponseWriter, r *http.Request) {
	http.NotFound(w, r)
}
