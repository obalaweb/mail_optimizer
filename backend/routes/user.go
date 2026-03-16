package routes

import (
	"github.com/gofiber/fiber/v2"
)

func RegisterUserRoutes(router fiber.Router) {
	router.Get("/me", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "user profile not implemented"})
	})
} 