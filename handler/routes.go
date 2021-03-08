package handler

import (
	"github.com/gorilla/mux"
)

func Routes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/ping", PingHandler())
	router.HandleFunc("/login", LoginHandler(deps.userService)).Methods("POST")
	router.HandleFunc("/assets/{asset_id}/maintenance", CreateMaintenanceHandler(deps.assetMaintenanceService)).Methods("POST")
	router.HandleFunc("/maintenance_activities/{id:[0-9]+}", DetailedMaintenanceActivityHandler(deps.assetMaintenanceService)).Methods("GET")
	router.HandleFunc("/users", ListUsersHandler(deps.userService)).Methods("GET")
	router.HandleFunc("/assets", ListAssetHandler(deps.assetService)).Methods("GET")
	return router

}
