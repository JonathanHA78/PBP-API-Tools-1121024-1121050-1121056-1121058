package goredis

import (
	"apitools/model"
	"context"
	"encoding/json"
	"log"
	"time"

	// "time"

	"github.com/redis/go-redis/v9"
)

func GetTasksRedis() []model.Task {
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

	var tasks []model.Task
	_ = json.Unmarshal([]byte(value), &tasks)

	return tasks
}

func SetTasksRedis(users []model.User) {
	converted, err := json.Marshal(users)
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

	err = client.Set(ctx, "users", converted, 5*time.Second).Err()
	if err != nil {
		log.Println("Set Error")
		log.Println(err)
		return
	} else {
		log.Println("Cache set")
	}
}

func DeleteUsersCache() error {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	var ctx = context.Background()

	err := client.Del(ctx, "users").Err()
	if err != nil {
		log.Println("Delete Error")
		log.Println(err)
		return err
	} else {
		log.Println("Cache deleted")
		return nil
	}
}
