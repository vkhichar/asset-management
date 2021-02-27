package handler

import (
	"fmt"
	"net/http"
)

func PingHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("pong"))
		if err != nil {
			fmt.Printf("handler: error while responding to ping: %s", err.Error())
		}
		return
	}
}
