package controllers

import (
	"fmt"
	"log"
	"net/http"
)

func TestConnection(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()
	fmt.Println("Sukses bos!")
	SendSuccessResponse(w, "Sukses Connect!")
}

func GetTaskListDaily(userID int) []Task {
	db := connect()
	defer db.Close()
	query := "SELECT id,user_id,title,description,due_date,completed FROM tasks where DATE(due_date)= CURDATE()AND completed=0 AND user_id =? ORDER BY due_date ASC"
	rows, err := db.Query(query, userID)
	if err != nil {
		log.Fatal(err)
	}
	var task Task
	var tasks []Task
	for rows.Next() {
		if err := rows.Scan(&task.Id, &task.UserId, &task.Title, &task.Description, &task.DueTime, &task.Complete); err != nil {
			log.Fatal(err)
		} else {
			tasks = append(tasks, task)
		}
	}
	return tasks
}

func InsertTask(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	user_id := r.Form.Get("user_id")
	title := r.Form.Get("title")
	desc := r.Form.Get("description")
	dueTime := r.Form.Get("due_time")
	completed := 0

	_, errQuery := db.Exec("INSERT INTO tasks(user_id ,title, description, due_date, completed ) values (?,?,?,?,?)",
		user_id, title, desc, dueTime, completed,
	)
	if errQuery == nil {
		SendSuccessResponse(w, "berhasil")
	} else {
		SendErrorResponse(w, "unkown error")
		fmt.Println(errQuery)
	}
}
