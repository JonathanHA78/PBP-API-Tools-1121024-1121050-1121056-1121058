package controllers

import (
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

}
