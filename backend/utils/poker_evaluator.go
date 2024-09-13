package utils

import (
	"sort"
)

type Card struct {
	Rank int
	Suit string
}

type Hand struct {
	Cards    []Card
	Name     string
	Rank     int
	Sequence []int // New field to store the sequence of the hand
}

const (
	HighCard = iota
	Pair
	TwoPair
	ThreeOfAKind
	Straight
	Flush
	FullHouse
	FourOfAKind
	StraightFlush
	RoyalFlush
)

func EvaluateHand(playerCards []Card, tableCards []Card) Hand {
	allCards := append(playerCards, tableCards...)
	bestHand := findBestHand(allCards)

	// Ensure that the best hand includes at least one of the player's cards
	if !handIncludesPlayerCard(bestHand, playerCards) {
		bestHand = findBestHandWithPlayerCard(allCards, playerCards)
	}

	return bestHand
}

func FindOtherBestHands(tableCards []Card) []Hand {
	allPossibleHands := generateAllPossibleHands(tableCards)
	sort.Slice(allPossibleHands, func(i, j int) bool {
		return allPossibleHands[i].Rank > allPossibleHands[j].Rank
	})

	// Return top 3 unique hands
	var uniqueHands []Hand
	for _, hand := range allPossibleHands {
		if len(uniqueHands) == 0 || hand.Name != uniqueHands[len(uniqueHands)-1].Name {
			uniqueHands = append(uniqueHands, hand)
			if len(uniqueHands) == 3 {
				break
			}
		}
	}

	return uniqueHands
}

func findBestHand(cards []Card) Hand {
	// Sort cards by rank in descending order
	sort.Slice(cards, func(i, j int) bool {
		return cards[i].Rank > cards[j].Rank
	})

	// Check for each hand type from highest to lowest
	if hand := checkRoyalFlush(cards); hand.Name != "" {
		return hand
	}
	if hand := checkStraightFlush(cards); hand.Name != "" {
		return hand
	}
	if hand := checkFourOfAKind(cards); hand.Name != "" {
		return hand
	}
	if hand := checkFullHouse(cards); hand.Name != "" {
		return hand
	}
	if hand := checkFlush(cards); hand.Name != "" {
		return hand
	}
	if hand := checkStraight(cards); hand.Name != "" {
		return hand
	}
	if hand := checkThreeOfAKind(cards); hand.Name != "" {
		return hand
	}
	if hand := checkTwoPair(cards); hand.Name != "" {
		return hand
	}
	if hand := checkPair(cards); hand.Name != "" {
		return hand
	}

	// If no other hand, return high card
	return Hand{Cards: cards[:5], Name: "High Card", Rank: HighCard, Sequence: getSequence(cards[:5])}
}

func handIncludesPlayerCard(hand Hand, playerCards []Card) bool {
	for _, card := range hand.Cards {
		for _, playerCard := range playerCards {
			if card == playerCard {
				return true
			}
		}
	}
	return false
}

func findBestHandWithPlayerCard(allCards []Card, playerCards []Card) Hand {
	var bestHand Hand
	for _, playerCard := range playerCards {
		remainingCards := removeCard(allCards, playerCard)
		for i := 0; i < len(remainingCards); i++ {
			handCards := append([]Card{playerCard}, remainingCards[:i]...)
			handCards = append(handCards, remainingCards[i+1:]...)
			hand := findBestHand(handCards[:5])
			if hand.Rank > bestHand.Rank {
				bestHand = hand
			}
		}
	}
	return bestHand
}

func generateAllPossibleHands(tableCards []Card) []Hand {
	var allHands []Hand
	for i := 0; i < len(tableCards)-1; i++ {
		for j := i + 1; j < len(tableCards); j++ {
			playerCards := []Card{tableCards[i], tableCards[j]}
			hand := EvaluateHand(playerCards, tableCards)
			allHands = append(allHands, hand)
		}
	}
	return allHands
}

func checkRoyalFlush(cards []Card) Hand {
	for _, suit := range []string{"Hearts", "Diamonds", "Clubs", "Spades"} {
		royalFlush := []int{14, 13, 12, 11, 10}
		count := 0
		var flushCards []Card
		for _, card := range cards {
			if card.Suit == suit && contains(royalFlush, card.Rank) {
				count++
				flushCards = append(flushCards, card)
			}
		}
		if count == 5 {
			return Hand{Cards: flushCards, Name: "Royal Flush", Rank: RoyalFlush, Sequence: royalFlush}
		}
	}
	return Hand{}
}

func checkStraightFlush(cards []Card) Hand {
	for _, suit := range []string{"Hearts", "Diamonds", "Clubs", "Spades"} {
		suitCards := filterBySuit(cards, suit)
		if len(suitCards) >= 5 {
			if straight := checkStraight(suitCards); straight.Name != "" {
				return Hand{Cards: straight.Cards, Name: "Straight Flush", Rank: StraightFlush, Sequence: straight.Sequence}
			}
		}
	}
	return Hand{}
}

func checkFourOfAKind(cards []Card) Hand {
	for i := 0; i <= len(cards)-4; i++ {
		if cards[i].Rank == cards[i+1].Rank && cards[i].Rank == cards[i+2].Rank && cards[i].Rank == cards[i+3].Rank {
			fourOfAKind := cards[i : i+4]
			kicker := getKicker(cards, fourOfAKind[0].Rank)
			sequence := []int{fourOfAKind[0].Rank, fourOfAKind[0].Rank, fourOfAKind[0].Rank, fourOfAKind[0].Rank, kicker.Rank}
			return Hand{Cards: append(fourOfAKind, kicker), Name: "Four of a Kind", Rank: FourOfAKind, Sequence: sequence}
		}
	}
	return Hand{}
}

func checkFullHouse(cards []Card) Hand {
	if three := checkThreeOfAKind(cards); three.Name != "" {
		remainingCards := removeCards(cards, three.Cards[:3])
		if pair := checkPair(remainingCards); pair.Name != "" {
			sequence := []int{three.Cards[0].Rank, three.Cards[0].Rank, three.Cards[0].Rank, pair.Cards[0].Rank, pair.Cards[0].Rank}
			return Hand{Cards: append(three.Cards[:3], pair.Cards[:2]...), Name: "Full House", Rank: FullHouse, Sequence: sequence}
		}
	}
	return Hand{}
}

func checkFlush(cards []Card) Hand {
	for _, suit := range []string{"Hearts", "Diamonds", "Clubs", "Spades"} {
		suitCards := filterBySuit(cards, suit)
		if len(suitCards) >= 5 {
			sequence := getSequence(suitCards[:5])
			return Hand{Cards: suitCards[:5], Name: "Flush", Rank: Flush, Sequence: sequence}
		}
	}
	return Hand{}
}

func checkStraight(cards []Card) Hand {
	for i := 0; i <= len(cards)-5; i++ {
		if cards[i].Rank-cards[i+4].Rank == 4 {
			sequence := getSequence(cards[i : i+5])
			return Hand{Cards: cards[i : i+5], Name: "Straight", Rank: Straight, Sequence: sequence}
		}
	}
	// Check for Ace-low straight
	if cards[0].Rank == 14 && cards[len(cards)-4].Rank == 5 && cards[len(cards)-3].Rank == 4 && cards[len(cards)-2].Rank == 3 && cards[len(cards)-1].Rank == 2 {
		aceLowStraight := append(cards[len(cards)-4:], cards[0])
		sequence := []int{5, 4, 3, 2, 1} // Ace is considered 1 in this case
		return Hand{Cards: aceLowStraight, Name: "Straight", Rank: Straight, Sequence: sequence}
	}
	return Hand{}
}

func checkThreeOfAKind(cards []Card) Hand {
	for i := 0; i <= len(cards)-3; i++ {
		if cards[i].Rank == cards[i+1].Rank && cards[i].Rank == cards[i+2].Rank {
			threeOfAKind := cards[i : i+3]
			kickers := getKickers(cards, threeOfAKind[0].Rank, 2)
			sequence := []int{threeOfAKind[0].Rank, threeOfAKind[0].Rank, threeOfAKind[0].Rank, kickers[0].Rank, kickers[1].Rank}
			return Hand{Cards: append(threeOfAKind, kickers...), Name: "Three of a Kind", Rank: ThreeOfAKind, Sequence: sequence}
		}
	}
	return Hand{}
}

func checkTwoPair(cards []Card) Hand {
	if pair1 := checkPair(cards); pair1.Name != "" {
		remainingCards := removeCards(cards, pair1.Cards[:2])
		if pair2 := checkPair(remainingCards); pair2.Name != "" {
			kicker := getKicker(cards, pair1.Cards[0].Rank, pair2.Cards[0].Rank)
			sequence := []int{pair1.Cards[0].Rank, pair1.Cards[0].Rank, pair2.Cards[0].Rank, pair2.Cards[0].Rank, kicker.Rank}
			return Hand{Cards: append(append(pair1.Cards[:2], pair2.Cards[:2]...), kicker), Name: "Two Pair", Rank: TwoPair, Sequence: sequence}
		}
	}
	return Hand{}
}

func checkPair(cards []Card) Hand {
	for i := 0; i <= len(cards)-2; i++ {
		if cards[i].Rank == cards[i+1].Rank {
			pair := cards[i : i+2]
			kickers := getKickers(cards, pair[0].Rank, 3)
			sequence := []int{pair[0].Rank, pair[0].Rank, kickers[0].Rank, kickers[1].Rank, kickers[2].Rank}
			return Hand{Cards: append(pair, kickers...), Name: "Pair", Rank: Pair, Sequence: sequence}
		}
	}
	return Hand{}
}

func filterBySuit(cards []Card, suit string) []Card {
	var filtered []Card
	for _, card := range cards {
		if card.Suit == suit {
			filtered = append(filtered, card)
		}
	}
	return filtered
}

func getKicker(cards []Card, excludeRanks ...int) Card {
	for _, card := range cards {
		if !contains(excludeRanks, card.Rank) {
			return card
		}
	}
	return Card{}
}

func getKickers(cards []Card, excludeRank int, count int) []Card {
	var kickers []Card
	for _, card := range cards {
		if card.Rank != excludeRank {
			kickers = append(kickers, card)
			if len(kickers) == count {
				break
			}
		}
	}
	return kickers
}

func removeCards(cards []Card, toRemove []Card) []Card {
	var result []Card
	for _, card := range cards {
		if !containsCard(toRemove, card) {
			result = append(result, card)
		}
	}
	return result
}

func removeCard(cards []Card, toRemove Card) []Card {
	var result []Card
	for _, card := range cards {
		if card != toRemove {
			result = append(result, card)
		}
	}
	return result
}

func contains(slice []int, val int) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

func containsCard(slice []Card, card Card) bool {
	for _, item := range slice {
		if item == card {
			return true
		}
	}
	return false
}

func getSequence(cards []Card) []int {
	sequence := make([]int, len(cards))
	for i, card := range cards {
		sequence[i] = card.Rank
	}
	return sequence
}
