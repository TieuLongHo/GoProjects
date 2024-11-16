package main

import (
	. "TodoWebApp/internal/controller"
	"TodoWebApp/internal/models"
	"TodoWebApp/internal/repository"
	"net/http"
)

func main() {
	repository.CreateTodo(models.TodoRequest{Description: "test"})
	repository.CreateTodo(models.TodoRequest{Description: "moin"})
	repository.CreateTodo(models.TodoRequest{Description: "build web app"})
	repository.CreateTodo(models.TodoRequest{Description: "deploy app"})
	repository.CreateTodo(models.TodoRequest{Description: "write documentation"})
	repository.DeleteTodo(3)

	http.HandleFunc("GET /todos", GetAllTodos)
	http.HandleFunc("GET /todos/{id}", GetTodo)
	http.HandleFunc("POST /todos", CreateTodo)
	http.HandleFunc("PUT /todos/{id}", CheckTodo)
	http.HandleFunc("DELETE /todos/{id}", DeleteTodo)
	http.ListenAndServe(":8080", nil)
}
