package controllers

import (
	"apitools/model"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/go-co-op/gocron"
)

func CreateCronExpression(t time.Time) string {
	year := strconv.Itoa(t.Year())
	month := strconv.Itoa(int(t.Month()))
	day := strconv.Itoa(t.Day())
	hour := strconv.Itoa(t.Hour())
	minute := strconv.Itoa(t.Minute())

	cronExpression := strings.Join([]string{minute, hour, day, month, "*", year}, " ")

	return cronExpression
}

// func CreateSchedule(t time.Time, todo func()) {
// 	layout := "2006-01-02 15:04:05"
// 	timeSchedule := t.Format(layout)
// 	// d := t.Sub(time.Now())
// 	// fmt.Println(t)
// 	// fmt.Println(time.Now())
// 	// fmt.Println(d)
// 	cron := CreateCronExpression(t)
// 	fmt.Println(cron)
// 	fmt.Println(timeSchedule)
// 	s := gocron.NewScheduler(time.Local)
// 	s.At(t.String()).Do(todo)
// 	// s.Every(d).Do(todo)
// 	// s.At(time.Now()).Do(todo)
// 	// _, err := s.Every(1).Minute().
// 	// if err != nil {
// 	// 	log.Fatal(err)
// 	// }
// 	s.StartAsync()
// }

func SendDailyEmail() {
	s := gocron.NewScheduler(time.Local)
	s.Every(1).Day().At("01:48").Do(SendEmailToAll)
	s.StartAsync()
}

func SendReminderMail(userId string, task model.Task) {
	db := connect()
	defer db.Close()
	query := "select name, email from users where id = ?"
	var name string
	var email string
	var tasks []model.Task
	tasks = append(tasks, task)
	errQuery := db.QueryRow(query, userId).Scan(&name, &email)
	fmt.Println("Akan dikrimkan ke ", email, " pada ", task.DueTime)
	if errQuery == nil {
		// content := GenerateEmail(1, name, tasks)
		content2 := GenerateEmail(3, name, tasks)

		SendEmail(content2, email)
		// go CreateSchedule(task.DueTime, func() {
		// 	fmt.Println("email pengingat terkirim")
		// 	SendEmail(content, email)
		// })
	} else {
		log.Fatal(errQuery)
	}
}
