package handler

import (
	"github.com/gorilla/mux"
)

func Routes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/ping", PingHandler())
	router.HandleFunc("/login", LoginHandler(deps.userService)).Methods("POST")

	router.HandleFunc("/assets/{asset_id}/maintenance", CreateMaintenanceHandler(deps.assetMaintenanceService)).Methods("POST")
	router.HandleFunc("/maintenance_activities/{id}", DetailedMaintenanceActivityHandler(deps.assetMaintenanceService)).Methods("GET")

	router.HandleFunc("/users", ListUsersHandler(deps.userService)).Methods("GET")
	router.HandleFunc("/users/{id}", UpdateUsersHandler(deps.userService)).Methods("PUT")
	router.HandleFunc("/users/{id}", DeleteUserHandler(deps.userService)).Methods("DELETE")

	router.HandleFunc("/assets", ListAssetHandler(deps.assetService)).Methods("GET")
	router.HandleFunc("/assets/{Id}", UpdateAssetHandler(deps.assetService)).Methods("PUT")
	router.HandleFunc("/assets/{Id}", DeleteAssetHandler(deps.assetService)).Methods("DELETE")
	router.HandleFunc("/users/{id}", DeleteUserHandler(deps.userService)).Methods("DELETE")

	// maintenance activities
	router.HandleFunc("/maintenance_activities/{id:[0-9]+}", AuthenticationHandler(deps.tokenService,
		DeleteMaintenanceActivityHandler(deps.assetMaintenanceService), true)).Methods("DELETE")
	router.HandleFunc("/assets/{asset_id}/maintenance", AuthenticationHandler(deps.tokenService,
		ListMaintenanceActivitiesByAsserId(deps.assetMaintenanceService), true)).Methods("GET")
	router.HandleFunc("/maintenance_activities/{id:[0-9]+}", AuthenticationHandler(deps.tokenService,
		UpdateMaintenanceActivity(deps.assetMaintenanceService), true)).Methods("PUT")

	router.HandleFunc("/token/verify", VerifyTokenHandler(deps.tokenService)).Methods("GET")
	router.HandleFunc("/token/generate", GenerateTokenHandler(deps.tokenService)).Methods("POST")

	return router
}
