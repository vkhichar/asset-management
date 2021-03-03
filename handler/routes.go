package handler

import "net/http"

func Routes() {
	http.Handle("/ping", PingHandler())
	http.Handle("/login", LoginHandler(deps.userService))
	http.Handle("/users", ListUsersHandler(deps.userService))
}
