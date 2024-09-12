package setup

import (
	"github.com/Sandy143toce/poker-evaluator/backend/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// Health check route
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	// Poker evaluation route
	app.Post("/evaluate", handlers.EvaluatePokerHand)

	// Recent game results route
	app.Get("/recent-results", handlers.GetRecentGameResults)

	// Add more routes here as needed
}
