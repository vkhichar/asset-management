package main

import (
	"fmt"
	"net/http"

	"github.com/vkhichar/asset-management/config"
	"github.com/vkhichar/asset-management/handler"
	"github.com/vkhichar/asset-management/repository"
)

func main() {
	err := config.Init()
	if err != nil {
		fmt.Printf("main: error while initialising config: %s", err.Error())
		return
	}

	// initialise db connection
	repository.InitDB()
	handler.InitDependencies()

	err = http.ListenAndServe(":"+config.GetAppPort(), handler.Routes())

	if err != nil {
		fmt.Printf("main: error while starting server: %s", err.Error())
		return
	}
}
