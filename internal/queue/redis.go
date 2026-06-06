package queue

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

var (
	Ctx = context.Background()
	RDB *redis.Client
)

func Connect() {

	RDB = redis.NewClient(
		&redis.Options{
			Addr: "localhost:6379",
		},
	)

	_, err := RDB.Ping(Ctx).Result()

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to Redis")
}
