package queue

import (
	"context"
	"log"
	"os"

	"github.com/redis/go-redis/v9"
)

var (
	Ctx = context.Background()
	RDB *redis.Client
)

func Connect() {

	addr := os.Getenv("REDIS_ADDR")

	if addr == "" {
		addr = "localhost:6379"
	}

	RDB = redis.NewClient(&redis.Options{
		Addr:     addr,
		Username: os.Getenv("default"),
		Password: os.Getenv("sxdAvMNm3ioaomJ266rm1lPyHqEljGVA"),
	})

	_, err := RDB.Ping(Ctx).Result()

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to Redis")
}
