package handler

import (
	"com/willredington/chat-api/redis"
	"com/willredington/chat-api/service"
	"log"

	"github.com/gofiber/fiber/v2"
)

func Trigger(c *fiber.Ctx) error {

	userId := "user-1"

	isWaiting, err := service.IsInWaitQueue(redis.Client, userId)
	if err != nil {
		return err
	}

	if !isWaiting {

		otherUserId, err := service.PopWaitQueue(redis.Client)
		if err != nil {
			return err
		}

		if otherUserId == "" {
			service.PushToWaitQueue(redis.Client, []string{userId})
			log.Println("nothing in the queue")
			return c.SendStatus(fiber.StatusNoContent)
		}

		// should never happen
		if otherUserId == userId {
			log.Println("expected user IDs to not be equal")
			return fiber.ErrInternalServerError
		}

		if _, err := service.CreateRoom(redis.Client, []string{userId, otherUserId}); err != nil {
			log.Println("failed to create room")
			return fiber.ErrInternalServerError
		}

		return c.SendStatus(fiber.StatusCreated)
	}

	log.Printf("user ID: %s is already in the waiting queue", userId)

	return c.SendStatus(fiber.StatusNoContent)

}
