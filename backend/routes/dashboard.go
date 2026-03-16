package routes

import (
	"github.com/gofiber/fiber/v2"
)

func RegisterDashboardRoutes(router fiber.Router) {
	router.Get("/summary", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "dashboard summary not implemented"})
	})
} 