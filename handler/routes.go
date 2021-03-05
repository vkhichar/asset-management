package handler

import (
	"github.com/gorilla/mux"
)

func Routes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/ping", PingHandler())
	router.HandleFunc("/login", LoginHandler(deps.userService)).Methods("POST")
	router.HandleFunc("/assets/{assetId}/maintenance", CreateMaintenanceHandler(deps.assetMaintenanceService)).Methods("POST")

	return router

}
