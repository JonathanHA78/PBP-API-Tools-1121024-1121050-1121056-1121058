package controllers

import (
	"apitools/model"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/go-co-op/gocron"
)

func createCronExpression(t time.Time) string {
	year := strconv.Itoa(t.Year())
	month := strconv.Itoa(int(t.Month()))
	day := strconv.Itoa(t.Day())
	hour := strconv.Itoa(t.Hour())
	minute := strconv.Itoa(t.Minute())

	cronExpression := strings.Join([]string{minute, hour, day, month, "*", year}, " ")

	return cronExpression
}

func CreateSchedule(t time.Time, todo func()) {
	s := gocron.NewScheduler(time.UTC)
	s.Cron(createCronExpression(t)).Do(todo)
	s.StartAsync()
}

func SendDailyEmail() {
	s := gocron.NewScheduler(time.UTC)
	s.Cron("0 0 6 * * *").Do(SendEmailToAll)
	s.StartAsync()
}

func SendReminderMail(userId string, task model.Task) {
	db := connect()
	defer db.Close()
	query := "select name, email from users where user_id = ?"
	var name string
	var email string
	var tasks []model.Task
	tasks = append(tasks, task)
	errQuery := db.QueryRow(query, userId).Scan(&name, &email)
	if errQuery == nil {
		content := GenerateEmail(1, name, tasks)
		CreateSchedule(task.DueTime, func() { SendEmail(content, email) })
	} else {
		log.Fatal(errQuery)
	}
}