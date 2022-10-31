package main

import (
	"com/willredington/chat-api/handler"
	"com/willredington/chat-api/middleware"

	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(".env"); err != nil {
		panic("Error loading .env file")
	}

	app := fiber.New()

	app.Use(adaptor.HTTPMiddleware(middleware.EnsureValidToken()))

	app.Get("/", handler.Hello)

	app.Listen(":9000")
}
