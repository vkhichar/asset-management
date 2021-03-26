package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type configs struct {
	jwtConfig      JwtConfig
	appPort        int
	dbConfig       DBConfig
	eventApiConfig EventApiConfig
}

type DBConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	Name     string
}

type JwtConfig struct {
	TokenExpiry int
	Secret      string
}

type EventApiConfig struct {
	Host    string
	Timeout int // in seconds
}

var config configs

const DEFAULT_TOKEN_EXPIRY = 5

func Init() error {
	portStr := os.Getenv("APP_PORT")
	port, err := strconv.Atoi(portStr)
	if err != nil {
		fmt.Printf("config: couldn't covert app_port from string to int: %s", err.Error())
		port = 9000
	}

	config.appPort = port
	config.dbConfig = initDBConfig()
	config.jwtConfig = initJwtConfig()
	config.eventApiConfig = initEventApiConfig()
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

func GetDBConfig() DBConfig {
	return config.dbConfig
}

func GetEventServiceUrl() string {
	return config.eventApiConfig.Host
}

func GetEventApiTimeout() int {
	return config.eventApiConfig.Timeout
}

func initJwtConfig() JwtConfig {
	tokenExpiry, err := strconv.Atoi(os.Getenv("TOKEN_EXPIRY"))
	if err != nil {
		fmt.Printf("config: couldn't read environment variable for token expiry: %s", err.Error())
		tokenExpiry = DEFAULT_TOKEN_EXPIRY
	}
	return JwtConfig{
		TokenExpiry: tokenExpiry,
		Secret:      os.Getenv("JWT_SECRET"),
	}
}

func GetJwtConfig() JwtConfig {
	return config.jwtConfig
}

func initEventApiConfig() EventApiConfig {
	eventApiConfig := EventApiConfig{}
	eventApiConfig.Host = os.Getenv("EVENT_SERVICE_URL")
	if strings.TrimSpace(eventApiConfig.Host) == "" {
		panic("config: missing EVENT_SERVICE_URL")
	}

	timeout, err := strconv.Atoi(os.Getenv("EVENT_API_TIMEOUT"))
	if err != nil {
		fmt.Println("config: Invalid timeout value: ", err)
		timeout = 3 // in seconds
	}
	eventApiConfig.Timeout = timeout
	return eventApiConfig
}
