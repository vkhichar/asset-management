package main

import (
	"fmt"
	"github.com/vkhichar/asset-management/config"
	"github.com/vkhichar/asset-management/handler"
	"github.com/vkhichar/asset-management/repository"
	"net/http"
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
	handler.Routes()

	err = http.ListenAndServe(":"+config.GetAppPort(), nil)
	if err != nil {
		fmt.Printf("main: error while starting server: %s", err.Error())
		return
	}
}
