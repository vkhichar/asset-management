package service

import (
	"os"

	"github.com/vkhichar/asset-management/config"
)

func InitEnv() {
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_PORT", "5432")
	os.Setenv("DB_USERNAME", "postgres")
	os.Setenv("DB_PASSWORD", "abc123")
	os.Setenv("EVENT_SERVICE_URL", "http://34.70.86.33:9035")
	os.Setenv("JWT_SECRET", "secret")
	os.Setenv("APP_PORT", "9000")
	os.Setenv("TOKEN_EXPIRY", "5")
	os.Setenv("EVENT_API_TIMEOUT", "3")
	config.Init()
}
