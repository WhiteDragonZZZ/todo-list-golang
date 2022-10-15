package main

import (
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"todo-list-golang"
	"todo-list-golang/pkg/handler"
	"todo-list-golang/pkg/repository"
	"todo-list-golang/pkg/service"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	db, err := repository.NewPostgres(repository.Config{
		Host:     "localhost",
		Port:     "5432",
		Username: "postgres",
		Password: "admin",
		DBName:   "todo-list",
		SSLMode:  "disable",
	})

	if err != nil {
		logrus.Fatalf("error occured while running db server: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	srv := new(todo_list_golang.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error occured while running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
