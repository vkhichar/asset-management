package service

import (
	"os"

	"github.com/vkhichar/asset-management/config"
)

func InitEnv() {
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_PORT", "5432")
	os.Setenv("DB_USERNAME", "postgres")
	os.Setenv("DB_PASSWORD", "12345")
	os.Setenv("EVENT_SERVICE_URL", "http://34.70.86.33:9035")
	os.Setenv("JWT_SECRET", "secret")
	config.Init()
}
