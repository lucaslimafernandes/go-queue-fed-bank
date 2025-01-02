package handler

import (
	"context"
	"log"
	"os"

	"github.com/go-redis/redis/v8"
)

var RDB *redis.Client

// RedisConnect establishes a connection to a Redis server using the provided configuration.
// It returns a pointer to the redis.Client object representing the connection and logs an error if something goes wrong.
func RedisConn() {

	redisAddr := os.Getenv("REDIS_ADDR")
	redisPasswd := os.Getenv("REDIS_PASSWORD")

	RDB = redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: redisPasswd,
		DB:       0,
	})

	ctx := context.Background()

	_, err := RDB.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("failed to connect to Redis: %v\n", err)
	}

}
