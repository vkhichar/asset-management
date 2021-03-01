package main

import (
	"os"

	"github.com/vkhichar/asset-management/server"
)

func main() {
	key := os.Args[1]

	switch key {
	case "start":
		server.Start()
	case "seed":
		return server.Insert_data()
		// case "migrate":
		// 	return db.RunMigrations()
		// case "rollback":
		//  return db.RollbackMigrations(os.Args().Get(0))
	default:
		server.Start()
	}
}
