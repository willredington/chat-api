package router

import (
	"com/willredington/chat-api/config"
	"com/willredington/chat-api/handler"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

// SetupRoutes setup router api
func SetupRoutes(app *fiber.App, config *config.AppConfig) {
	// Middleware
	api := app.Group(fmt.Sprintf("/api/%s", config.ApiVersion))

	// room
	room := api.Group("/room")
	room.Get("/", handler.GetRoom)
	room.Post("/", handler.CreateRoom)

	// message
	message := api.Group("/message")
	message.Get("/", handler.GetMessages)
	message.Post("/", handler.SendMessage)

}
