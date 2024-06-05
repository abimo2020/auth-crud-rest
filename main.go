package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/Auth-CRUD-REST/config"
	"github.com/Auth-CRUD-REST/models/seeder"
	"github.com/Auth-CRUD-REST/routes"
	"github.com/Auth-CRUD-REST/services"
	"github.com/Auth-CRUD-REST/util"
	"github.com/go-playground/validator"
)

func init() {
	config.InitDB()
	config.InitialMigration()
	seeder.DBSeed(config.DB)

	config.RedisInit()

	config.InitRabbitMQ()
}

func main() {
	e := routes.New()
	e.Validator = &util.CustomValidator{
		Validator: validator.New(),
	}
	go services.ConsumeMessage()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		if err := e.Start(":8080"); err != nil {
			e.Logger.Info("Shutting down the server")
		}
	}()
	<-quit
}
