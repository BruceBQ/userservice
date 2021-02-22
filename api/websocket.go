package api

import (
	"net/http"

	"github.com/gorilla/websocket"
)

func (api *API) InitWebSocket() {
	api.BaseRoutes.ApiRoot.HandleFunc("/", connectWebSocket).Methods("GET")
}

func connectWebSocket(w http.ResponseWriter, r *http.Request) {
	upgrader := websocket.Upgrader{}

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {

	}

	ws.Close()
}
