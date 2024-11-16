package models

type Todo struct {
	Id          uint   `json:"id"`
	Description string `json:"description"`
	IsDone      bool   `json:"isDone"`
}

type TodoRequest struct {
	Description string `json:description`
}
