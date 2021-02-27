package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/vkhichar/asset-management/config"
)

var sqlxDB *sqlx.DB

func InitDB() {
	db, err := sqlx.Open("postgres", getConnectionString())
	if err != nil {
		panic(fmt.Sprintf("app: error while opening DB connection: %s", err.Error()))
	}

	if err = db.Ping(); err != nil {
		panic(fmt.Sprintf("app: error while opening DB connection: %s", err.Error()))
	}

	sqlxDB = db
}

func getConnectionString() string {
	dbConfig := config.GetDBConfig()
	return fmt.Sprintf("dbname=%s user=%s password='%s' host=%s port=%d sslmode=disable", dbConfig.Name, dbConfig.Username, dbConfig.Password, dbConfig.Host, dbConfig.Port)
}

func GetDB() *sqlx.DB {
	return sqlxDB
}
