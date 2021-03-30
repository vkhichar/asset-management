package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/urfave/cli"
	"github.com/vkhichar/asset-management/config"
	"github.com/vkhichar/asset-management/handler"
	"github.com/vkhichar/asset-management/repository"
)

func main() {

	err := config.Init()
	if err != nil {
		fmt.Printf("main: error while initialising config: %s", err.Error())
		return
	}
	cliApp := cli.NewApp()
	cliApp.Name = "asset_management"
	cliApp.Usage = "Details of Assets"
	cliApp.Commands = []cli.Command{
		{
			Name:  "startapp",
			Usage: "Start Server",
			Action: func(c *cli.Context) {
				StartApp()
			},
		},
		{
			Name:  "migrate",
			Usage: "Migrate DataBase Migrations",
			Action: func(c *cli.Context) {
				repository.RunMigrations()
			},
		},
		{
			Name:  "rollback",
			Usage: "RollBack DataBase Migrations",
			Action: func(c *cli.Context) {
				repository.RollBackMigrations()
			},
		},
	}

	err = cliApp.Run(os.Args)
	if err != nil {
		log.Printf("something went wrong", err.Error())
	}

}
func StartApp() {
	// initialise db connection
	repository.InitDB()
	handler.InitDependencies()

	srv := &http.Server{
		Addr:    ":" + config.GetAppPort(),
		Handler: handler.Routes(),
	}
	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		err := srv.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatalf("main: error while starting server: %s", err)
		}
	}()

	log.Print("Server Started")
	<-done
	log.Print("Server Stopped")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer func() {
		cancel()
	}()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown Failed:%+v", err)
	}
	log.Print("Server Exited Properly")
}
