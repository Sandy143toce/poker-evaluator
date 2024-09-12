package handlers

import (
	"fmt"

	"github.com/Sandy143toce/poker-evaluator/backend/database"
	"github.com/Sandy143toce/poker-evaluator/backend/models"
	"github.com/Sandy143toce/poker-evaluator/backend/utils"
	"github.com/gofiber/fiber/v2"
)

func EvaluatePokerHand(c *fiber.Ctx) error {
	var request models.PokerEvaluationRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Error: "Invalid request body",
		})
	}

	// Validate the request
	if len(request.PlayerCards) != 2 || len(request.TableCards) < 3 || len(request.TableCards) > 5 {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Error: "Invalid number of cards",
		})
	}

	// Convert request cards to utils.Card
	playerCards := convertToCards(request.PlayerCards)
	tableCards := convertToCards(request.TableCards)

	// Evaluate the hand
	hand := utils.EvaluateHand(playerCards, tableCards)

	// Prepare the response
	response := models.PokerEvaluationResponse{
		Hand:     hand.Name,
		HandRank: hand.Rank,
		Cards:    convertToStringCards(hand.Cards),
	}

	// Store the result in the database
	db := database.GetDB()
	err := database.StoreGameResult(db, response)
	if err != nil {
		// Log the error, but don't return it to the client
		fmt.Printf("Failed to store game result: %v\n", err)
	}

	// Update the cached recent game results
	redisClient := utils.GetRedisClient()
	recentResults, _ := database.GetRecentGameResults(db, 10)
	_ = utils.CacheRecentGameResults(redisClient, recentResults)

	return c.JSON(response)
}

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

func convertToCards(stringCards []string) []utils.Card {
	cards := make([]utils.Card, len(stringCards))
	for i, stringCard := range stringCards {
		cards[i] = parseCard(stringCard)
	}
	return cards
}

func convertToStringCards(cards []utils.Card) []string {
	stringCards := make([]string, len(cards))
	for i, card := range cards {
		stringCards[i] = formatCard(card)
	}
	return stringCards
}

func parseCard(stringCard string) utils.Card {
	rankMap := map[string]int{
		"2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9, "10": 10,
		"J": 11, "Q": 12, "K": 13, "A": 14,
	}
	suitMap := map[string]string{
		"H": "Hearts", "D": "Diamonds", "C": "Clubs", "S": "Spades",
	}

	rank := rankMap[stringCard[:len(stringCard)-1]]
	suit := suitMap[stringCard[len(stringCard)-1:]]

	return utils.Card{Rank: rank, Suit: suit}
}

func formatCard(card utils.Card) string {
	rankMap := map[int]string{
		2: "2", 3: "3", 4: "4", 5: "5", 6: "6", 7: "7", 8: "8", 9: "9", 10: "10",
		11: "J", 12: "Q", 13: "K", 14: "A",
	}
	suitMap := map[string]string{
		"Hearts": "H", "Diamonds": "D", "Clubs": "C", "Spades": "S",
	}

	return rankMap[card.Rank] + suitMap[card.Suit]
}
