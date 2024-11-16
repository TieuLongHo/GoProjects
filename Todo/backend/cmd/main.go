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

	mux := http.NewServeMux()

	mux.HandleFunc("GET /todos", GetAllTodos)
	mux.HandleFunc("GET /todos/{id}", GetTodo)
	mux.HandleFunc("POST /todos", CreateTodo)
	mux.HandleFunc("PUT /todos/{id}", CheckTodo)
	mux.HandleFunc("DELETE /todos/{id}", DeleteTodo)
	http.ListenAndServe(":8080", corsMiddleware(mux))
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
