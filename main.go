package main

import (
	"fmt"
	"net/http"

	"github.com/vkhichar/asset-management/config"
	"github.com/vkhichar/asset-management/handler"
	"github.com/vkhichar/asset-management/repository"

	"github.com/newrelic/go-agent/v3/newrelic"
)

func main() {

	app, _ := newrelic.NewApplication(
		newrelic.ConfigAppName("asset_management"),
		newrelic.ConfigLicense("3bbe15e00ed7cf0f2c378311e2d558f16812NRAL"),
		newrelic.ConfigDistributedTracerEnabled(true),
	)

	err := config.Init()
	if err != nil {
		fmt.Printf("main: error while initialising config: %s", err.Error())
		return
	}

	// initialise db connection
	repository.InitDB()
	handler.InitDependencies()
	err = http.ListenAndServe(":"+config.GetAppPort(), handler.Routes(app))

	if err != nil {
		fmt.Printf("main: error while starting server: %s", err.Error())
		return
	}
}
