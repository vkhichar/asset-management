package handler

import "net/http"

func Routes() {
	http.Handle("/ping", PingHandler(deps))
	http.Handle("/login", LoginHandler(deps))
}
