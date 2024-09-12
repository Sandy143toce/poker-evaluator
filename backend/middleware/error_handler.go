package middleware

import (
	"log"

	"github.com/Sandy143toce/poker-evaluator/backend/models"
	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(c *fiber.Ctx, err error) error {
	// Log the error
	log.Printf("Error: %v", err)

	// Default error message
	message := "An unexpected error occurred"
	status := fiber.StatusInternalServerError

	// Check if it's a Fiber error
	if e, ok := err.(*fiber.Error); ok {
		message = e.Message
		status = e.Code
	}

	// Return JSON response
	return c.Status(status).JSON(models.ErrorResponse{
		Error: message,
	})
}
