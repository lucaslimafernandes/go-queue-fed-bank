package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/lucaslimafernandes/go-queue-fed-bank/api/handler"
	"github.com/lucaslimafernandes/go-queue-fed-bank/api/router"
)

func init() {

	handler.RedisConn()

}

func main() {

	app := fiber.New()
	app.Use(cors.New())

	router.SetupRoutes(app)

	log.Fatal(app.Listen(os.Getenv("PORT")))

}
