package controller

import (
	"TodoWebApp/internal/models"
	"TodoWebApp/internal/repository"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func GetAllTodos(w http.ResponseWriter, req *http.Request) {
	fmt.Println("GET /todos called")
	todos := repository.GetAll()
	response, err := json.Marshal(todos)
	if err != nil {
		http.Error(w, "Failed to marshal JSON", http.StatusInternalServerError)
		log.Printf("Error marshaling JSON: %v, Request: %v\n", err, req.URL.Path)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
func GetTodo(w http.ResponseWriter, req *http.Request) {
	id, err := strconv.Atoi(req.PathValue("id"))
	if err != nil {
		http.Error(w, "Failed to parse ID", http.StatusBadRequest)
		log.Printf("Error parsing ID: %v, Request: %v\n", err, req.URL.Path)
		return
	}
	fmt.Printf("GET /todos/%v called\n", id)

	todo, err := repository.GetTodo(uint(id))
	if err != nil {
		http.Error(w, "Failed to find Todo", http.StatusNotFound)

		log.Printf("Error finding Todo: %v, Request: %v\n", err, req.URL.Path)
		return
	}
	response, err := json.Marshal(todo)
	if err != nil {
		http.Error(w, "Failed to marshal JSON", http.StatusInternalServerError)
		log.Printf("Error marshaling JSON: %v, Request: %v\n", err, req.URL.Path)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)

}

func CreateTodo(w http.ResponseWriter, req *http.Request) {
	fmt.Println("POST /todos called")
	var todoReq models.TodoRequest

	err := json.NewDecoder(req.Body).Decode(&todoReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	newTodo, _ := repository.CreateTodo(todoReq)
	w.WriteHeader(http.StatusCreated)
	resp, err := json.Marshal(newTodo)
	if err != nil {
		http.Error(w, "Failed to marshal JSON", http.StatusInternalServerError)
		log.Printf("Error marshaling JSON: %v, Request: %v\n", err, req.URL.Path)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)

}

func CheckTodo(w http.ResponseWriter, req *http.Request) {
	fmt.Println("PUT /todos called")
	var todoReq models.Todo

	err := json.NewDecoder(req.Body).Decode(&todoReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	todo, err := repository.UpdateTodo(todoReq)

	w.WriteHeader(http.StatusOK)
	resp, err := json.Marshal(todo)
	if err != nil {
		http.Error(w, "Failed to marshal JSON", http.StatusInternalServerError)
		log.Printf("Error marshaling JSON: %v, Request: %v\n", err, req.URL.Path)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)

}

func DeleteTodo(w http.ResponseWriter, req *http.Request) {
	id, err := strconv.Atoi(req.PathValue("id"))
	if err != nil {
		http.Error(w, "Failed to parse ID", http.StatusBadRequest)
		log.Printf("Error parsing ID: %v, Request: %v\n", err, req.URL.Path)
		return
	}
	fmt.Printf("DELETE /todos/%v called\n", id)

	err = repository.DeleteTodo(uint(id))
	if err != nil {
		http.Error(w, "Failed to find Todo", http.StatusNotFound)
		log.Printf("Error finding Todo: %v, Request: %v\n", err, req.URL.Path)
		return
	}

	w.WriteHeader(http.StatusOK)

}
