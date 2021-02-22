package api

import "net/http"

func (api *API) InitHealth() {
	api.BaseRoutes.Health.HandleFunc("", getHealth).Methods("GET")
}

func getHealth(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}
