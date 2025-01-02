package handler

import (
	"github.com/gofiber/fiber/v2"
)

func In(c *fiber.Ctx) error {

	// TODO

	return c.JSON(fiber.Map{"status": "success", "message": "Inserted on queue", "data": nil})

}

func Ping(c *fiber.Ctx) error {

	pong, err := RDB.Ping(c.Context()).Result()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": err, "data": nil})
	}
	return c.JSON(fiber.Map{"status": "success", "message": "Pong", "data": pong})

}
