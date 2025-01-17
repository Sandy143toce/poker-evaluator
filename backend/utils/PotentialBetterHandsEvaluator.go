package utils

import (
	"sort"

	"github.com/Sandy143toce/poker-evaluator/backend/models"
)

func GetPotentialBetterHands(playerCards []Card, tableCards []Card, currentBestHand Hand) []models.PotentialBetterHand {
	var potentialHands []models.PotentialBetterHand

	// Get all possible ranks and suits excluding player's cards
	// Generate all possible cards and remove the ones already in use
	allCards := generateAvailableCards()
	usedCards := append(playerCards, tableCards...)
	availableCards := removeUsedCards(allCards, usedCards)

	// For each pair of available cards
	for i := 0; i < len(availableCards); i++ {
		for j := i + 1; j < len(availableCards); j++ {
			testCards := []Card{availableCards[i], availableCards[j]}
			possibleHand := GetBestHand(testCards, tableCards)

			// If this hand is better than player's current hand
			if possibleHand.Rank < currentBestHand.Rank {
				betterHand := models.PotentialBetterHand{
					Name:        possibleHand.Name,
					HandRank:    possibleHand.Rank,
					NeededCards: []string{formatCard(testCards[0]), formatCard(testCards[1])},
				}

				// Check if we already have this type of hand
				if !containsSimilarHand(potentialHands, betterHand) {
					potentialHands = append(potentialHands, betterHand)
				}
			}
		}
	}

	// Sort by hand rank (lower is better in our ranking system)
	sortPotentialHands(potentialHands)

	// Return top 3 or all if less than 3
	if len(potentialHands) > 3 {
		return potentialHands[:3]
	}
	return potentialHands
}

func generateAvailableCards() []Card {
	var cards []Card
	suits := []string{"Hearts", "Diamonds", "Clubs", "Spades"}
	ranks := []int{2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}

	for _, suit := range suits {
		for _, rank := range ranks {
			cards = append(cards, Card{Rank: rank, Suit: suit})
		}
	}
	return cards
}

func removeUsedCards(cards []Card, usedCards []Card) []Card {
	var availableCards []Card
	for _, card := range cards {
		if !containsCard(usedCards, card) {
			availableCards = append(availableCards, card)
		}
	}
	return availableCards
}

func containsSimilarHand(hands []models.PotentialBetterHand, newHand models.PotentialBetterHand) bool {
	for _, hand := range hands {
		if hand.Name == newHand.Name {
			return true
		}
	}
	return false
}

func sortPotentialHands(hands []models.PotentialBetterHand) {
	sort.Slice(hands, func(i, j int) bool {
		return hands[i].HandRank < hands[j].HandRank
	})
}

func formatCard(card Card) string {
	rankMap := map[int]string{
		2: "2", 3: "3", 4: "4", 5: "5", 6: "6", 7: "7", 8: "8", 9: "9", 10: "10",
		11: "J", 12: "Q", 13: "K", 14: "A",
	}
	suitMap := map[string]string{
		"Hearts": "H", "Diamonds": "D", "Clubs": "C", "Spades": "S",
	}
	return rankMap[card.Rank] + suitMap[card.Suit]
}
