package controllers

import (
	"apitools/model"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

func TestConnection(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()
	fmt.Println("Sukses bos!")
	SendSuccessResponse(w, "Sukses Connect!")
}
func GetAllUsers() []model.User {
	db := connect()
	defer db.Close()
	query := "SELECT id,name,email FROM users"
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	var user model.User
	var users []model.User
	for rows.Next() {
		if err := rows.Scan(&user.Id, &user.Name, &user.Email); err != nil {
			log.Fatal(err)
		} else {
			users = append(users, user)
		}
	}
	return users
}

func GetTaskListDaily(userID int) []model.Task {
	var tasks []model.Task
	key := "user" + strconv.Itoa(userID)
	tasks = GetUserTasksRedis(key)
	if tasks == nil {
		db := connect()
		defer db.Close()
		query := "SELECT id,user_id,title,description,due_date,completed FROM tasks where DATE(due_date)= CURDATE()AND completed=0 AND user_id =? ORDER BY due_date ASC"
		rows, err := db.Query(query, userID)
		if err != nil {
			log.Fatal(err)
		}
		var task model.Task
		for rows.Next() {
			if err := rows.Scan(&task.Id, &task.UserId, &task.Title, &task.Description, &task.DueTime, &task.Complete); err != nil {
				log.Fatal(err)
			} else {
				tasks = append(tasks, task)
			}
		}
		SetUserTasksRedis(tasks, key)
	}

	return tasks
}

func StringToTime(dateStr string) time.Time {
	layout := "2006-01-02 15:04:05"
	date, err := time.Parse(layout, dateStr)
	if err != nil {
		log.Fatal("Error parsing date string:", err)
	}
	utcPlus7 := time.FixedZone("UTC+7", 7*60*60)
	return date.In(utcPlus7)
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
		userId, _ := strconv.Atoi(user_id)
		tasks := GetTaskListDaily(userId)
		SetUserTasksRedis(tasks, "user"+user_id)
		var task model.Task
		task.Title = title
		task.Description = desc
		task.DueTime = StringToTime(dueTime)
		SendReminderMail(user_id, task)
		SendSuccessResponse(w, "berhasil")
	} else {
		SendErrorResponse(w, "unkown error")
		fmt.Println(errQuery)
	}
}
