package controllers

import (
	"apitools/model"
	"context"
	"encoding/json"
	"log"
	"time"

	// "time"

	"github.com/redis/go-redis/v9"
)

func GetUserTasksRedis(key string) []model.Task {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	var ctx = context.Background()

	value, err := client.Get(ctx, key).Result()
	if err != nil {
		log.Println("Get Error")
		log.Println(err)
		return nil
	}

	var tasks []model.Task
	_ = json.Unmarshal([]byte(value), &tasks)

	return tasks
}

func SetUserTasksRedis(tasks []model.Task, key string) {
	converted, err := json.Marshal(tasks)
	if err != nil {
		log.Println(err)
		return
	}

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	var ctx = context.Background()

	err = client.Set(ctx, key, converted, 5*time.Second).Err()
	if err != nil {
		log.Println("Set Error")
		log.Println(err)
		return
	} else {
		log.Println("Cache set")
	}
}

func DeleteUserTasksCache(key string) error {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	var ctx = context.Background()

	err := client.Del(ctx, key).Err()
	if err != nil {
		log.Println("Delete Error")
		log.Println(err)
		return err
	} else {
		log.Println("Cache deleted")
		return nil
	}
}
