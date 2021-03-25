package handler

import (
	"github.com/gorilla/mux"
)

func Routes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/ping", PingHandler())
	router.HandleFunc("/login", LoginHandler(deps.userService)).Methods("POST")

	router.HandleFunc("/users", AuthenticationHandler(deps.tokenService,
		CreateUserHandler(deps.userService), true)).Methods("POST")
	router.HandleFunc("/users", AuthenticationHandler(deps.tokenService, ListUsersHandler(deps.userService), true)).Methods("GET")
	router.HandleFunc("/profile", UserAuthenticationHandler(deps.tokenService, GetUserByIDHandler(deps.userService))).Methods("GET")
	router.HandleFunc("/profile", UserAuthenticationHandler(deps.tokenService,
		UpdateUsersHandler(deps.userService))).Methods("PUT")
	router.HandleFunc("/profile", UserAuthenticationHandler(deps.tokenService,
		DeleteUserHandler(deps.userService))).Methods("DELETE")

	router.HandleFunc("/assets", CreateAssetHandler(deps.assetService)).Methods("POST")
	router.HandleFunc("/assets", ListAssetHandler(deps.assetService)).Methods("GET")
	router.HandleFunc("/assets/{id}", GetAssetHandler(deps.assetService)).Methods("GET")
	router.HandleFunc("/assets/{Id}", UpdateAssetHandler(deps.assetService)).Methods("PUT")
	router.HandleFunc("/assets/{Id}", DeleteAssetHandler(deps.assetService)).Methods("DELETE")

	router.HandleFunc("/assets/{asset_id}/maintenance", CreateMaintenanceHandler(deps.assetMaintenanceService)).Methods("POST")
	router.HandleFunc("/maintenance_activities/{id}", DetailedMaintenanceActivityHandler(deps.assetMaintenanceService)).Methods("GET")

	// maintenance activities
	router.HandleFunc("/maintenance_activities/{id:[0-9]+}", AuthenticationHandler(deps.tokenService,
		DeleteMaintenanceActivityHandler(deps.assetMaintenanceService), true)).Methods("DELETE")
	router.HandleFunc("/assets/{asset_id}/maintenance", AuthenticationHandler(deps.tokenService,
		ListMaintenanceActivitiesByAsserId(deps.assetMaintenanceService), true)).Methods("GET")
	router.HandleFunc("/maintenance_activities/{id:[0-9]+}", AuthenticationHandler(deps.tokenService,
		UpdateMaintenanceActivity(deps.assetMaintenanceService), true)).Methods("PUT")

	router.HandleFunc("/token/verify", VerifyTokenHandler(deps.tokenService)).Methods("GET")
	router.HandleFunc("/token/generate", GenerateTokenHandler(deps.tokenService)).Methods("POST")

	router.HandleFunc("/assets/{asset_id}/allocate/{user_id}", AuthenticationHandler(deps.tokenService,
		CreateAssetAllocationHandler(deps.assetAllocationService), true)).Methods("POST")

	router.HandleFunc("/assets/{asset_id}/deallocate", AuthenticationHandler(deps.tokenService,
		AssetDeAllocationHandler(deps.assetAllocationService), true)).Methods("DELETE")
	// router.HandleFunc("/assets/{asset_id}/deallocate", AssetDeAllocationHandler(deps.assetAllocationService)).Methods("DELETE")
	return router
}
