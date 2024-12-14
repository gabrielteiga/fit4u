package usecase

import "github.com/gofiber/fiber/v2"

func Ping(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"success": true,
		"message": "pong",
	})
}
