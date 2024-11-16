package repository

import (
	. "TodoWebApp/internal/models"
	"fmt"
	"strconv"
	"unsafe"
)

var todos []Todo = make([]Todo, 0)

func GetAll() []Todo {
	return append([]Todo{}, todos...)
}
func GetTodo(id uint) (Todo, error) {

	todo, err := findTodo(id)
	return *todo, err
}
func CreateTodo(t Todo) error {
	todos = append(todos, t)
	return nil
}
func ToggleTodo(id uint) error {
	todo, err := findTodo(id)
	if err != nil {
		return err
	}
	todo.IsDone = !todo.IsDone
	return nil
}
func DeleteTodo(id uint) error {
	todosStartAddress := fmt.Sprintf("%p", &todos[0])
	todosStartAddressDecimal, _ := strconv.ParseInt(todosStartAddress[2:], 16, 64)

	todo, _ := findTodo(id)
	todoAdress := fmt.Sprintf("%p", todo)
	todoAddressDecimal, _ := strconv.ParseInt(todoAdress[2:], 16, 64)
	todoByteSize := unsafe.Sizeof(Todo{})
	todoIndex := (todoAddressDecimal - todosStartAddressDecimal) / int64(todoByteSize)
	todos = append(todos[:todoIndex], todos[todoIndex+1:]...)
	fmt.Println(todos)
	return nil
}
func findTodo(id uint) (*Todo, error) {
	for i, t := range todos {
		if t.Id == id {
			return &todos[i], nil
		}
	}
	return nil, fmt.Errorf("can't find todo")
}
