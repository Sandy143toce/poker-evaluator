package main

import (
	"log"
	"os"

	"github.com/Sandy143toce/poker-evaluator/backend/middleware"
	"github.com/Sandy143toce/poker-evaluator/backend/setup"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env if it exists (for local development)
	godotenv.Load() // Ignore error as file may not exist in production

	// Initialize database connection
	// db, err := database.InitDB()
	// if err != nil {
	// 	log.Fatalf("Failed to connect to database: %v", err)
	// }
	// defer db.Close()

	// Initialize Redis connection only if Redis config is present
	// if os.Getenv("REDIS_HOST") != "" {
	// 	redisClient := utils.InitRedis()
	// 	if redisClient != nil {
	// 		defer redisClient.Close()
	// 	}
	// } else {
	// 	log.Println("Redis configuration not found, skipping Redis initialization")
	// }

	// Create Fiber app
	app := fiber.New(fiber.Config{
		ErrorHandler: middleware.ErrorHandler,
	})

	// Middleware
	app.Use(logger.New(logger.Config{
		Format: "[${time}] ${status} - ${method} ${path}\n",
	}))

	// Configure CORS for both local development and production
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*", // In production, you might want to restrict this to your frontend domain
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	// Setup routes
	setup.SetupRoutes(app)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	if err := app.Listen(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
