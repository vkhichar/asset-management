package main

import (
	"fmt"
	"net/http"

	"github.com/afex/hystrix-go/hystrix"
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
	hystrix.ConfigureCommand("create_event", hystrix.CommandConfig{
		Timeout:                1000,
		MaxConcurrentRequests:  300,
		RequestVolumeThreshold: 3,
		SleepWindow:            2500,
		ErrorPercentThreshold:  20,
	})
	// initialise db connection
	repository.InitDB()
	handler.InitDependencies()
	err = http.ListenAndServe(":"+config.GetAppPort(), handler.Routes())

	if err != nil {
		fmt.Printf("main: error while starting server: %s", err.Error())
		return
	}
}
