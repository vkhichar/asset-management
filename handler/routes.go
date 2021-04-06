package handler

import (
	"github.com/gorilla/mux"
	"github.com/newrelic/go-agent/v3/newrelic"
)

func Routes(app *newrelic.Application) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc(newrelic.WrapHandleFunc(app, "/ping", PingHandler()))
	router.HandleFunc(newrelic.WrapHandleFunc(app, "/login", LoginHandler(deps.userService))).Methods("POST")

	router.HandleFunc(newrelic.WrapHandleFunc(app, "/users", AuthenticationHandler(deps.tokenService,
		CreateUserHandler(deps.userService), true))).Methods("POST")
	router.HandleFunc(newrelic.WrapHandleFunc(app, "/users", ListUsersHandler(deps.userService))).Methods("GET")
	router.HandleFunc(newrelic.WrapHandleFunc(app, "/users/{id}", GetUserByIDHandler(deps.userService))).Methods("GET")
	router.HandleFunc(newrelic.WrapHandleFunc(app, "/users/{id}", UpdateUsersHandler(deps.userService))).Methods("PUT")
	router.HandleFunc(newrelic.WrapHandleFunc(app, "/users/{id}", DeleteUserHandler(deps.userService))).Methods("DELETE")

	router.HandleFunc(newrelic.WrapHandleFunc(app, "/assets", CreateAssetHandler(deps.assetService))).Methods("POST")
	router.HandleFunc(newrelic.WrapHandleFunc(app, "/assets", ListAssetHandler(deps.assetService))).Methods("GET")
	router.HandleFunc(newrelic.WrapHandleFunc(app, "/assets/{id}", GetAssetHandler(deps.assetService))).Methods("GET")
	router.HandleFunc(newrelic.WrapHandleFunc(app, "/assets/{Id}", UpdateAssetHandler(deps.assetService))).Methods("PUT")
	router.HandleFunc(newrelic.WrapHandleFunc(app, "/assets/{Id}", DeleteAssetHandler(deps.assetService))).Methods("DELETE")

	router.HandleFunc(newrelic.WrapHandleFunc(app, "/assets/{asset_id}/maintenance", CreateMaintenanceHandler(deps.assetMaintenanceService))).Methods("POST")
	router.HandleFunc(newrelic.WrapHandleFunc(app, "/maintenance_activities/{id}", DetailedMaintenanceActivityHandler(deps.assetMaintenanceService))).Methods("GET")

	// maintenance activities
	router.HandleFunc(newrelic.WrapHandleFunc(app, "/maintenance_activities/{id:[0-9]+}", AuthenticationHandler(deps.tokenService,
		DeleteMaintenanceActivityHandler(deps.assetMaintenanceService), true))).Methods("DELETE")
	router.HandleFunc(newrelic.WrapHandleFunc(app, "/assets/{asset_id}/maintenance", AuthenticationHandler(deps.tokenService,
		ListMaintenanceActivitiesByAsserId(deps.assetMaintenanceService), true))).Methods("GET")
	router.HandleFunc(newrelic.WrapHandleFunc(app, "/maintenance_activities/{id:[0-9]+}", AuthenticationHandler(deps.tokenService,
		UpdateMaintenanceActivity(deps.assetMaintenanceService), true))).Methods("PUT")

	router.HandleFunc(newrelic.WrapHandleFunc(app, "/token/verify", VerifyTokenHandler(deps.tokenService))).Methods("GET")
	router.HandleFunc(newrelic.WrapHandleFunc(app, "/token/generate", GenerateTokenHandler(deps.tokenService))).Methods("POST")

	router.HandleFunc(newrelic.WrapHandleFunc(app, "/assets/{asset_id}/allocate/{user_id}", AuthenticationHandler(deps.tokenService,
		CreateAssetAllocationHandler(deps.assetAllocationService), true))).Methods("POST")

	router.HandleFunc(newrelic.WrapHandleFunc(app, "/assets/{asset_id}/deallocate", AuthenticationHandler(deps.tokenService,
		AssetDeAllocationHandler(deps.assetAllocationService), true))).Methods("DELETE")
	// router.HandleFunc("/assets/{asset_id}/deallocate", AssetDeAllocationHandler(deps.assetAllocationService)).Methods("DELETE")
	return router
}
