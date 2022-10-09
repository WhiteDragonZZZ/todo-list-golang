package main

import (
	"log"
	todo_list_golang "todo-list-golang"
	"todo-list-golang/pkg/handler"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(todo_list_golang.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
