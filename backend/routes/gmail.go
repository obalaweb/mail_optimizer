package routes

import (
	"github.com/gofiber/fiber/v2"
)

func RegisterGmailRoutes(router fiber.Router) {
	router.Get("/emails", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "fetch emails not implemented"})
	})
	router.Post("/emails/actions", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "email actions not implemented"})
	})
	router.Get("/labels", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "list labels not implemented"})
	})
	router.Post("/labels", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "create label not implemented"})
	})
	router.Put("/labels/:id", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "update label not implemented"})
	})
	router.Delete("/labels/:id", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "delete label not implemented"})
	})
} 