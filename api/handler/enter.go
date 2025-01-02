package handler

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
)

func In(c *fiber.Ctx) error {

	// TODO

	return c.JSON(fiber.Map{"status": "success", "message": "Inserted on queue", "data": nil})

}

func Ping(c *fiber.Ctx) error {

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "suasenha",
		DB:       0,
	})

	ctx := context.Background()

	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": err, "data": nil})
	}
	return c.JSON(fiber.Map{"status": "success", "message": "Pong", "data": pong})

}
