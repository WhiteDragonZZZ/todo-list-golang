package main

import (
	"log"
	todo_list_golang "todo-list-golang"
	"todo-list-golang/pkg/handler"
	"todo-list-golang/pkg/repository"
	"todo-list-golang/pkg/service"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	srv := new(todo_list_golang.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
