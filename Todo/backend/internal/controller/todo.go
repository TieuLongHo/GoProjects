package controller

import (
	"TodoWebApp/internal/repository"
	"net/http"
)

func GetAllTodos(w http.ResponseWriter, req *http.Request) {
	repository.GetAll()

}
func GetTodo(w http.ResponseWriter, req *http.Request) {

}

func CreateTodo(w http.ResponseWriter, req *http.Request) {

}

func CheckTodo(w http.ResponseWriter, req *http.Request) {

}

func DeleteTodo(w http.ResponseWriter, req *http.Request) {

}
