package main

import (
	. "TodoWebApp/internal/controller"
	"TodoWebApp/internal/models"
	"TodoWebApp/internal/repository"
	"net/http"
)

func main() {
	repository.CreateTodo(models.Todo{Id: 0, Description: "test", IsDone: true})
	repository.CreateTodo(models.Todo{Id: 1, Description: "moin", IsDone: false})
	repository.CreateTodo(models.Todo{Id: 2, Description: "build web app", IsDone: false})
	repository.CreateTodo(models.Todo{Id: 3, Description: "deploy app", IsDone: true})
	repository.CreateTodo(models.Todo{Id: 4, Description: "write documentation", IsDone: false})
	repository.DeleteTodo(3)

	http.HandleFunc("GET /todos", GetAllTodos)
	http.HandleFunc("GET /todos/{id}", GetTodo)
	http.HandleFunc("POST /todos", CreateTodo)
	http.HandleFunc("PATCH /todos/{id}", CheckTodo)
	http.HandleFunc("DELETE /todos/{id}", DeleteTodo)
	http.ListenAndServe(":8080", nil)
}
