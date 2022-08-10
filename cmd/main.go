package main

import (
	"log"

	"github.com/GaponovAlexey/practic-web"
	"github.com/GaponovAlexey/practic-web/pkg/handler"
	"github.com/GaponovAlexey/practic-web/pkg/repository"
	"github.com/GaponovAlexey/practic-web/pkg/service"
)

func main() {

	repos := repository.NewRepository()
	service := service.NewService(repos)
	handler := handler.NewHandler(service)

	//start
	srv := new(todo.Server)
	if err := srv.Run("3000", handler.InitRoutes()); err != nil {
		log.Fatal(err)
	}
}
