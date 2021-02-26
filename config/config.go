package config

import (
	"os"
	"strconv"
)

type configs struct {
	appPort int
}

var config configs

func Init() error {
	portStr := os.Getenv("APP_PORT")
	port, err := strconv.Atoi(portStr)
	if err != nil {
		port = 9000
	}

	config.appPort = port

	return nil
}

func GetAppPort() string {
	return strconv.Itoa(config.appPort)
}
