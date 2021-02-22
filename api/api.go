package api

import (
	"net/http"

	"github.com/gorilla/mux"

	"userservice/app"
	"userservice/model"
	"userservice/web"
)

type Routes struct {
	Root    *mux.Router
	ApiRoot *mux.Router // 'api/v1'

	Auth *mux.Router // 'api/v1/author'

	Users *mux.Router // 'api/v1/users'
	User  *mux.Router // 'api/v1/users/{user_id:[A-Za-z0-9]+}'

	Social *mux.Router // 'api/v1/social'

	Permission *mux.Router // 'api/v1/permission'

	Role *mux.Router // 'api/v1/roles'

	Health *mux.Router // 'api/v1/health'
}

type API struct {
	GetGlobalAppOptions app.AppOptionCreator
	BaseRoutes          *Routes
}

func Init(globalOptionsFunc app.AppOptionCreator, root *mux.Router) *API {
	api := &API{
		GetGlobalAppOptions: globalOptionsFunc,
		BaseRoutes:          &Routes{},
	}

	api.BaseRoutes.Root = root
	api.BaseRoutes.ApiRoot = root.PathPrefix(model.API_URL_SUFFIX).Subrouter()

	api.BaseRoutes.Auth = api.BaseRoutes.ApiRoot.PathPrefix("/authentication").Subrouter()

	api.BaseRoutes.Users = api.BaseRoutes.ApiRoot.PathPrefix("/users").Subrouter()
	api.BaseRoutes.User = api.BaseRoutes.ApiRoot.PathPrefix("/users/{user_id:[a-f0-9]+}").Subrouter()

	api.BaseRoutes.Social = api.BaseRoutes.ApiRoot.PathPrefix("/social").Subrouter()

	api.BaseRoutes.Permission = api.BaseRoutes.ApiRoot.PathPrefix("/permissions").Subrouter()

	api.BaseRoutes.Role = api.BaseRoutes.ApiRoot.PathPrefix("/roles").Subrouter()

	api.BaseRoutes.Health = api.BaseRoutes.ApiRoot.PathPrefix("/health").Subrouter()

	api.InitUser()
	api.InitSocical()
	api.InitHealth()
	api.InitAuthorization()
	api.InitPermission()
	api.InitRole()

	root.Handle("/api/v1/{anything:.*}", http.HandlerFunc(api.Handle404))
	return api
}

func (api *API) Handle404(w http.ResponseWriter, r *http.Request) {
	web.Handle404(w, r)
}
