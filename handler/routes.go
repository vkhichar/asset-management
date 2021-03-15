package handler

import (
	"github.com/gorilla/mux"
)

func Routes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/ping", PingHandler())
	router.HandleFunc("/login", LoginHandler(deps.userService)).Methods("POST")
	router.HandleFunc("/users", ListUsersHandler(deps.userService)).Methods("GET")
	router.HandleFunc("/assets", ListAssetHandler(deps.assetService)).Methods("GET")
	router.HandleFunc("/assets/{Id}", UpdateAssetHandler(deps.assetService)).Methods("PUT")
	router.HandleFunc("/assets/delete/{Id}", DeleteAssetHandler(deps.assetService)).Methods("DELETE")
	return router
}
