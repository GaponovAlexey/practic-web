package main

import (
	"os"

	"github.com/GaponovAlexey/practic-web"
	"github.com/GaponovAlexey/practic-web/pkg/handler"
	"github.com/GaponovAlexey/practic-web/pkg/repository"
	"github.com/GaponovAlexey/practic-web/pkg/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

func main() {
	//env
	err := godotenv.Load()
	if err != nil {
		logrus.Fatal("Error loading .env file")
	}
	//config LogrusJson
	logrus.SetFormatter(new(logrus.JSONFormatter))

	//db
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     os.Getenv("HOST"),
		Port:     os.Getenv("PORT"),
		Username: os.Getenv("BASENAME"),
		Password: os.Getenv("DB_PASSWORD"),
		DBname:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("SSLMODE"),
	})
	if err != nil {
		logrus.Fatal("failed to init", err)
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
		logrus.Fatal(err)
	}
}
