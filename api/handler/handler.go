package handler

import "github.com/gofiber/fiber/v2"

func In(c *fiber.Ctx) error {

	// TODO

	return c.JSON(fiber.Map{"status": "success", "message": "Inserted on queue", "data": nil})

}
