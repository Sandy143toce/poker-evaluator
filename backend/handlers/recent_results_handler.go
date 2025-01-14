package handlers

import (
	"github.com/Sandy143toce/poker-evaluator/backend/database"
	"github.com/Sandy143toce/poker-evaluator/backend/models"
	"github.com/Sandy143toce/poker-evaluator/backend/utils"
	"github.com/gofiber/fiber/v2"
)

func GetRecentGameResults(c *fiber.Ctx) error {
	redisClient := utils.GetRedisClient()
	cachedResults, err := utils.GetCachedRecentGameResults(redisClient)
	if err == nil && cachedResults != nil {
		return c.JSON(cachedResults)
	}

	// If cache miss or error, fetch from database
	db := database.GetDB()
	results, err := database.GetRecentGameResults(db, 10)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{
			Error: "Failed to fetch recent game results",
		})
	}

	// Update cache
	_ = utils.CacheRecentGameResults(redisClient, results)

	return c.JSON(results)
}
