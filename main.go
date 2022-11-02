package main

import (
	"com/willredington/chat-api/config"
	"com/willredington/chat-api/middleware"
	"com/willredington/chat-api/redis"
	"com/willredington/chat-api/router"
	"log"

	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {

	app := fiber.New()
	config, err := config.New()
	if err != nil {
		log.Fatal("error loading config", err)
	}

	redis.ConnectRedis(config)

	if config.IsDev {
		log.Println("adding logging middleware...")
		app.Use(logger.New())
	}

	app.Use(cors.New())
	app.Use(adaptor.HTTPMiddleware(middleware.EnsureValidToken()))

	router.SetupRoutes(app, config)

	app.Listen(":9000")
}
