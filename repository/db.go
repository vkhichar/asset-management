package repository

import (
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/vkhichar/asset-management/config"
)

const (
	dbDriver      = "postgres"
	migrationPath = "./migrations"
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
func RunMigrations() {
	db, err := sqlx.Open(dbDriver, getConnectionString())
	if err != nil {
		panic(fmt.Sprintf("app: error while opening DB connection: %s", err.Error()))
	}

	if err = db.Ping(); err != nil {
		panic(fmt.Sprintf("app: error while opening DB connection: %s", err.Error()))
	}
	sqlDB := db.DB
	driver, err := postgres.WithInstance(sqlDB, &postgres.Config{})
	if err != nil {
		return
	}

	m, err := migrate.NewWithDatabaseInstance(getMigrationPath(), dbDriver, driver)
	if err != nil {
		return
	}
	err = m.Up()

	if err == migrate.ErrNoChange || err != nil {
		err = nil
		// fmt.Printf("ERROR %s", err.Error())
		return
	}
	log.Printf("Database Successfully Migrated")

}
func RollBackMigrations() {
	db, err := sqlx.Open("postgres", getConnectionString())
	if err != nil {
		panic(fmt.Sprintf("app: error while opening DB connection: %s", err.Error()))
	}

	if err = db.Ping(); err != nil {
		panic(fmt.Sprintf("app: error while opening DB connection: %s", err.Error()))
	}
	sqlDB := db.DB
	driver, err := postgres.WithInstance(sqlDB, &postgres.Config{})
	if err != nil {
		return
	}

	m, err := migrate.NewWithDatabaseInstance(getMigrationPath(), dbDriver, driver)
	if err != nil {
		return
	}
	err = m.Down()
	if err == migrate.ErrNoChange || err != nil {
		// err = nil
		// fmt.Printf("ERROR %s", err.Error())
		return
	}
	log.Printf("Database Successfully RollBacked")

}

func getConnectionString() string {
	dbConfig := config.GetDBConfig()
	return fmt.Sprintf("dbname=%s user=%s password='%s' host=%s port=%d sslmode=disable", dbConfig.Name, dbConfig.Username, dbConfig.Password, dbConfig.Host, dbConfig.Port)
}

func GetDB() *sqlx.DB {
	return sqlxDB
}
func getMigrationPath() string {
	return fmt.Sprintf("file://%s", migrationPath)
}
