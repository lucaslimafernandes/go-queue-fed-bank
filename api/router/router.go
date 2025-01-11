package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lucaslimafernandes/go-queue-fed-bank/api/handler"
)

func SetupRoutes(app *fiber.App) {

	// HealthCheck

	h := app.Group("/")
	h.Get("", handler.Ping)

	// queues
	q := app.Group("/queues")
	q.Post("/", handler.In)
	q.Get("/", handler.Show)

}
