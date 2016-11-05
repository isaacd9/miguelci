package queue

import (
	"fmt"
	"gopkg.in/redis.v5"
)

type QueueManager struct {
	Client *redis.Client
}

var (
	Manager QueueManager
)

func Connect() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	Manager.Client = client
}
