package utils

import (
	"sort"
)

type Card struct {
	Rank int
	Suit string
}

type Hand struct {
	Cards []Card
	Name  string
	Rank  int
}

const (
	RoyalFlush = iota
	StraightFlush
	FourOfAKind
	FullHouse
	Flush
	Straight
	ThreeOfAKind
	TwoPair
	Pair
	HighCard
)

func GetBestHand(playerCards []Card, tableCards []Card) Hand {
	allCards := append(playerCards, tableCards...)

	// Sort cards by rank in descending order
	sort.Slice(allCards, func(i, j int) bool {
		return allCards[i].Rank > allCards[j].Rank
	})

	// Check for each hand type from highest to lowest
	if hand := checkRoyalFlush(allCards); hand.Name != "" {
		return hand
	}
	if hand := checkStraightFlush(allCards); hand.Name != "" {
		return hand
	}
	if hand := checkFourOfAKind(allCards); hand.Name != "" {
		return hand
	}
	if hand := checkFullHouse(allCards); hand.Name != "" {
		return hand
	}
	if hand := checkFlush(allCards); hand.Name != "" {
		return hand
	}
	if hand := checkStraight(allCards); hand.Name != "" {
		return hand
	}
	if hand := checkThreeOfAKind(allCards); hand.Name != "" {
		return hand
	}
	if hand := checkTwoPair(allCards); hand.Name != "" {
		return hand
	}
	if hand := checkPair(allCards); hand.Name != "" {
		return hand
	}

	// If no other hand, return high card
	return Hand{Cards: allCards[:5], Name: "High Card", Rank: HighCard}
}

func checkRoyalFlush(cards []Card) Hand {
	for _, suit := range []string{"Hearts", "Diamonds", "Clubs", "Spades"} {
		royalFlush := []int{14, 13, 12, 11, 10}
		count := 0
		for _, card := range cards {
			if card.Suit == suit && contains(royalFlush, card.Rank) {
				count++
			}
		}
		if count == 5 {
			return Hand{Cards: cards[:5], Name: "Royal Flush", Rank: RoyalFlush}
		}
	}
	return Hand{}
}

func checkStraightFlush(cards []Card) Hand {
	for _, suit := range []string{"Hearts", "Diamonds", "Clubs", "Spades"} {
		suitCards := filterBySuit(cards, suit)
		if len(suitCards) >= 5 {
			if straight := checkStraight(suitCards); straight.Name != "" {
				return Hand{Cards: straight.Cards, Name: "Straight Flush", Rank: StraightFlush}
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
			return Hand{Cards: append(fourOfAKind, kicker), Name: "Four of a Kind", Rank: FourOfAKind}
		}
	}
	return Hand{}
}

func checkFullHouse(cards []Card) Hand {
	if three := checkThreeOfAKind(cards); three.Name != "" {
		remainingCards := removeCards(cards, three.Cards[:3])
		if pair := checkPair(remainingCards); pair.Name != "" {
			return Hand{Cards: append(three.Cards[:3], pair.Cards[:2]...), Name: "Full House", Rank: FullHouse}
		}
	}
	return Hand{}
}

func checkFlush(cards []Card) Hand {
	for _, suit := range []string{"Hearts", "Diamonds", "Clubs", "Spades"} {
		suitCards := filterBySuit(cards, suit)
		if len(suitCards) >= 5 {
			return Hand{Cards: suitCards[:5], Name: "Flush", Rank: Flush}
		}
	}
	return Hand{}
}

func checkStraight(cards []Card) Hand {
	// Get unique ranks and sort them
	rankMap := make(map[int]Card)
	ranks := []int{}
	for _, card := range cards {
		if _, exists := rankMap[card.Rank]; !exists {
			rankMap[card.Rank] = card
			ranks = append(ranks, card.Rank)
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(ranks)))

	// Check for regular straight
	for i := 0; i <= len(ranks)-5; i++ {
		if ranks[i]-ranks[i+4] == 4 {
			// Found a straight - get the actual cards
			straightCards := make([]Card, 5)
			for j := 0; j < 5; j++ {
				straightCards[j] = rankMap[ranks[i+j]]
			}
			return Hand{Cards: straightCards, Name: "Straight", Rank: Straight}
		}
	}

	// Check for Ace-low straight
	hasAce := false
	has2to5 := true
	aceCard := Card{}
	lowStraightCards := make([]Card, 0, 5)

	for _, card := range cards {
		if card.Rank == 14 {
			hasAce = true
			aceCard = card
		}
	}

	if hasAce {
		for _, rank := range []int{2, 3, 4, 5} {
			if _, exists := rankMap[rank]; !exists {
				has2to5 = false
				break
			}
		}
		if has2to5 {
			for _, rank := range []int{2, 3, 4, 5} {
				lowStraightCards = append(lowStraightCards, rankMap[rank])
			}
			lowStraightCards = append(lowStraightCards, aceCard)
			return Hand{Cards: lowStraightCards, Name: "Straight", Rank: Straight}
		}
	}

	return Hand{}
}

func checkThreeOfAKind(cards []Card) Hand {
	for i := 0; i <= len(cards)-3; i++ {
		if cards[i].Rank == cards[i+1].Rank && cards[i].Rank == cards[i+2].Rank {
			threeOfAKind := cards[i : i+3]
			kickers := getKickers(cards, threeOfAKind[0].Rank, 2)
			return Hand{Cards: append(threeOfAKind, kickers...), Name: "Three of a Kind", Rank: ThreeOfAKind}
		}
	}
	return Hand{}
}

func checkTwoPair(cards []Card) Hand {
	if pair1 := checkPair(cards); pair1.Name != "" {
		remainingCards := removeCards(cards, pair1.Cards[:2])
		if pair2 := checkPair(remainingCards); pair2.Name != "" {
			kicker := getKicker(cards, pair1.Cards[0].Rank, pair2.Cards[0].Rank)
			return Hand{Cards: append(append(pair1.Cards[:2], pair2.Cards[:2]...), kicker), Name: "Two Pair", Rank: TwoPair}
		}
	}
	return Hand{}
}

func checkPair(cards []Card) Hand {
	for i := 0; i <= len(cards)-2; i++ {
		if cards[i].Rank == cards[i+1].Rank {
			pair := cards[i : i+2]
			kickers := getKickers(cards, pair[0].Rank, 3)
			return Hand{Cards: append(pair, kickers...), Name: "Pair", Rank: Pair}
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
