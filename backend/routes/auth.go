package routes

import (
	"github.com/gofiber/fiber/v2"
)

func RegisterAuthRoutes(router fiber.Router) {
	router.Post("/register", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "register not implemented"})
	})
	router.Post("/login", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "login not implemented"})
	})
	router.Get("/google", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "google oauth not implemented"})
	})
	router.Get("/google/callback", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "google callback not implemented"})
	})
	router.Post("/logout", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "logout not implemented"})
	})
} 