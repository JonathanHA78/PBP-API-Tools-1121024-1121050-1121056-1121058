package goredis

import (
	"apitools/controllers"
	"context"
	"encoding/json"
	"log"

	// "time"

	"github.com/redis/go-redis/v9"
)

func GetTasksRedis() []controllers.Task {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	var ctx = context.Background()

	value, err := client.Get(ctx, "users").Result()
	if err != nil {
		log.Println("Get Error")
		log.Println(err)
		return nil
	}

	var tasks []controllers.Task
	_ = json.Unmarshal([]byte(value), &tasks)

	return tasks
}
