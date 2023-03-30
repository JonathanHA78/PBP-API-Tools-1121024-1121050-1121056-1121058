package controllers

import "time"

type User struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Users struct {
	Users []User `json:"user_list"`
}

type Task struct {
	Id       int       `json:"id"`
	UserId   int       `json:"user_id"`
	Desc     string    `json:"description"`
	DueTime  time.Time `json:"due_time"`
	Complete int       `json:"complete"`
}

type Tasks struct {
	TaskList []Task `json:"task_list"`
}

type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type SuccessResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
