package handler

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Routes() {
	r := mux.NewRouter()
	http.Handle("/ping", PingHandler())
	http.Handle("/login", LoginHandler(deps.userService))
	r.HandleFunc("/assets/{assetid}/maintenance", CreateMaintenanceHandler(deps.userService))
}
