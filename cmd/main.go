package main

import (
	"log"

	"github.com/GaponovAlexey/practic-web"
	"github.com/GaponovAlexey/practic-web/pkg/handler"
	"github.com/GaponovAlexey/practic-web/pkg/repository"
	"github.com/GaponovAlexey/practic-web/pkg/service"
)

func main() {

	//db
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     "localhost",
		Port:     "5432",
		Username: "postgres",
		Password: "postgres",
		DBname:   "postgres",
		SSLMode:  "disable",
	})
	if err != nil {
		log.Fatal("failed to init", err)
	}

	//repos
	repos := repository.NewRepository(db)
	//service
	service := service.NewService(repos)
	//handler
	handler := handler.NewHandler(service)

	//start
	srv := new(todo.Server)
	if err := srv.Run("3000", handler.InitRoutes()); err != nil {
		log.Fatal(err)
	}
}
