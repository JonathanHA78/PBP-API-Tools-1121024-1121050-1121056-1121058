package controllers

import (
	"apitools/model"
)

func sendEmailToAll() {
	var users = GetAllUsers()
	for i := 0; i < len(users); i++ {
		var task []model.Task = GetTaskListDaily(users[i].Id)
		var content string = GenerateEmail(2, users[i], task)
		go SendEmail(content, users[i].Email)
	}
}
