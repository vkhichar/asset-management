package config

import (
	"fmt"
	"os"
	"strconv"
)

type configs struct {
	appPort   int
	eventPort int
	dbConfig  DBConfig
}

type DBConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	Name     string
}

var config configs

func Init() error {
	portStr := os.Getenv("APP_PORT")
	port, err := strconv.Atoi(portStr)
	if err != nil {
		fmt.Printf("config: couldn't covert app_port from string to int: %s", err.Error())
		port = 9000
	}
	eventportstr := os.Getenv("EVENT_PORT")
	eventport, err := strconv.Atoi(eventportstr)
	if err != nil {
		fmt.Printf("config: couldn't covert event_port from string to int: %s", err.Error())
		eventport = 9035
	}

	config.appPort = port
	config.eventPort = eventport
	config.dbConfig = initDBConfig()

	return nil
}

func initDBConfig() DBConfig {
	cfg := DBConfig{
		Host:     os.Getenv("DB_HOST"),
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Name:     os.Getenv("DB_NAME"),
	}

	portStr := os.Getenv("DB_PORT")
	port, err := strconv.Atoi(portStr)
	if err != nil {
		panic(fmt.Sprintf("config: couldn't read environment variable for db port: %s", err.Error()))
	}
	cfg.Port = port

	return cfg
}

func GetAppPort() string {
	return strconv.Itoa(config.appPort)
}
func GetEventPort() string {
	return strconv.Itoa(config.eventPort)
}

func GetDBConfig() DBConfig {
	return config.dbConfig
}
