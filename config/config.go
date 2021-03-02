package config

import (
	"fmt"
	"os"
	"strconv"
)

type configs struct {
	appPort   int
	dbConfig  DBConfig
	jwtConfig JwtConfig
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

var config configs

const DEFAULT_TOKEN_EXPIRY = 5 // in minutes

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
