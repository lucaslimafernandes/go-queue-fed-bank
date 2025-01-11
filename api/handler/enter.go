package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/lucaslimafernandes/go-queue-fed-bank/api/model"
)

func In(c *fiber.Ctx) error {

	userEntry := new(model.QueueEntry)
	err := c.BodyParser(userEntry)
	if err != nil {
		return c.Status(http.StatusInternalServerError).
			JSON(fiber.Map{"status": "InternalServerError", "message": err.Error(), "data": nil})
	}

	err = InsertUserQueue(userEntry.UserId)
	if err != nil {
		return c.Status(http.StatusInternalServerError).
			JSON(fiber.Map{"status": "InternalServerError", "message": err.Error(), "data": nil})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Inserted on queue", "data": nil})

}

func Show(c *fiber.Ctx) error {

	userId := new(model.QueueEntry)
	err := c.BodyParser(userId)
	if err != nil {
		return c.Status(http.StatusInternalServerError).
			JSON(fiber.Map{"status": "InternalServerError", "message": err.Error(), "data": nil})
	}

	user, err := GetUserByID(userId.UserId)
	if err != nil {
		return c.Status(http.StatusInternalServerError).
			JSON(fiber.Map{"status": "InternalServerError", "message": err.Error(), "data": nil})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Show", "data": user})

}

func Ping(c *fiber.Ctx) error {

	pong, err := RDB.Ping(c.Context()).Result()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": err, "data": nil})
	}
	return c.JSON(fiber.Map{"status": "success", "message": "Pong", "data": pong})

}
