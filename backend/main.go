package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/obalaweb/gmail_optimizer/routes"
)

func main() {
	app := fiber.New()

	// Enable CORS for all origins (customize as needed)
	app.Use(cors.New())

	api := app.Group("/api")

	routes.RegisterAuthRoutes(api.Group("/auth"))
	routes.RegisterUserRoutes(api.Group("/user"))
	routes.RegisterGmailRoutes(api.Group("/gmail"))
	routes.RegisterDashboardRoutes(api.Group("/dashboard"))

	log.Println("Server running on http://localhost:3000")
	if err := app.Listen(":3000"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
