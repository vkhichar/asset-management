package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type configs struct {
	appPort         int
	eventAppPort    int
	dbConfig        DBConfig
	eventServiceUrl string
	apiTimeout      int
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
	port, portErr := strconv.Atoi(portStr)
	if portErr != nil {
		fmt.Printf("config: couldn't covert app_port from string to int: %s", portErr.Error())
		port = 9000
	}

	eventPortStr := os.Getenv("EVENT_PORT")
	eventPort, err := strconv.Atoi(eventPortStr)
	if err != nil {
		fmt.Printf("config: couldn't covert app_port from string to int: %s", err.Error())
		port = 9035
	}

	config.appPort = port
	config.eventAppPort = eventPort
	config.dbConfig = initDBConfig()

	config.eventServiceUrl = os.Getenv("EVENT_SERVICE_URL")
	if strings.TrimSpace(config.eventServiceUrl) == "" {
		panic("config: missing EVENT_SERVICE_URL")
	}

	timeout, err := strconv.Atoi(os.Getenv("EVENT_API_TIMEOUT"))
	if err != nil {
		fmt.Println("config: Invalid timeout value: ", err)
		timeout = 3 // in seconds
	}
	config.apiTimeout = timeout
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

func GetEventAppPort() string {
	return strconv.Itoa(config.eventAppPort)
}

func GetDBConfig() DBConfig {
	return config.dbConfig
}

func GetEventServiceUrl() string {
	return config.eventServiceUrl
}

func GetEventApiTimeout() int {
	return config.apiTimeout
}
