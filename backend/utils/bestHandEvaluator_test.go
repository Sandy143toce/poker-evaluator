package utils

import (
	"testing"
)

func TestGetBestHand(t *testing.T) {
	tests := []struct {
		name        string
		playerCards []Card
		tableCards  []Card
		want        string
		wantRank    int
	}{
		{
			name: "Straight with pair in hand",
			playerCards: []Card{
				{Rank: 6, Suit: "Hearts"},
				{Rank: 6, Suit: "Diamonds"},
			},
			tableCards: []Card{
				{Rank: 4, Suit: "Clubs"},
				{Rank: 5, Suit: "Spades"},
				{Rank: 7, Suit: "Hearts"},
				{Rank: 8, Suit: "Hearts"},
				{Rank: 9, Suit: "Clubs"},
			},
			want:     "Straight",
			wantRank: Straight,
		},
		{
			name: "Royal Flush possible",
			playerCards: []Card{
				{Rank: 14, Suit: "Hearts"}, // Ace
				{Rank: 13, Suit: "Hearts"}, // King
			},
			tableCards: []Card{
				{Rank: 12, Suit: "Hearts"}, // Queen
				{Rank: 11, Suit: "Hearts"}, // Jack
				{Rank: 10, Suit: "Hearts"}, // 10
				{Rank: 9, Suit: "Hearts"},
				{Rank: 8, Suit: "Hearts"},
			},
			want:     "Royal Flush",
			wantRank: RoyalFlush,
		},
		{
			name: "Four of a Kind vs Full House",
			playerCards: []Card{
				{Rank: 10, Suit: "Hearts"},
				{Rank: 10, Suit: "Diamonds"},
			},
			tableCards: []Card{
				{Rank: 10, Suit: "Clubs"},
				{Rank: 10, Suit: "Spades"},
				{Rank: 9, Suit: "Hearts"},
				{Rank: 9, Suit: "Diamonds"},
				{Rank: 9, Suit: "Clubs"},
			},
			want:     "Four of a Kind",
			wantRank: FourOfAKind,
		},
		{
			name: "Ace-low Straight",
			playerCards: []Card{
				{Rank: 14, Suit: "Hearts"}, // Ace
				{Rank: 2, Suit: "Diamonds"},
			},
			tableCards: []Card{
				{Rank: 3, Suit: "Clubs"},
				{Rank: 4, Suit: "Spades"},
				{Rank: 5, Suit: "Hearts"},
				{Rank: 8, Suit: "Hearts"},
				{Rank: 9, Suit: "Clubs"},
			},
			want:     "Straight",
			wantRank: Straight,
		},
		{
			name: "Best straight when multiple possible",
			playerCards: []Card{
				{Rank: 6, Suit: "Hearts"},
				{Rank: 7, Suit: "Diamonds"},
			},
			tableCards: []Card{
				{Rank: 8, Suit: "Clubs"},
				{Rank: 9, Suit: "Spades"},
				{Rank: 10, Suit: "Hearts"},
				{Rank: 5, Suit: "Hearts"},
				{Rank: 4, Suit: "Clubs"},
			},
			want:     "Straight",
			wantRank: Straight,
		},
		{
			name: "Pair should not beat straight",
			playerCards: []Card{
				{Rank: 7, Suit: "Hearts"},
				{Rank: 7, Suit: "Diamonds"},
			},
			tableCards: []Card{
				{Rank: 8, Suit: "Clubs"},
				{Rank: 9, Suit: "Spades"},
				{Rank: 10, Suit: "Hearts"},
				{Rank: 6, Suit: "Hearts"},
				{Rank: 5, Suit: "Clubs"},
			},
			want:     "Straight",
			wantRank: Straight,
		},
		{
			name: "Three of a kind",
			playerCards: []Card{
				{Rank: 7, Suit: "Hearts"},
				{Rank: 7, Suit: "Diamonds"},
			},
			tableCards: []Card{
				{Rank: 7, Suit: "Clubs"},
				{Rank: 2, Suit: "Spades"},
				{Rank: 3, Suit: "Hearts"},
				{Rank: 4, Suit: "Hearts"},
				{Rank: 5, Suit: "Clubs"},
			},
			want:     "Three of a Kind",
			wantRank: ThreeOfAKind,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetBestHand(tt.playerCards, tt.tableCards)
			if got.Name != tt.want {
				t.Errorf("EvaluateHand() = %v, want %v", got.Name, tt.want)
			}
			if got.Rank != tt.wantRank {
				t.Errorf("EvaluateHand() rank = %v, want rank %v", got.Rank, tt.wantRank)
			}
		})
	}
}

// Helper function to create string cards for testing
func createCards(cards []string) []Card {
	result := make([]Card, len(cards))
	for i, card := range cards {
		rank := 0
		switch card[0] {
		case 'A':
			rank = 14
		case 'K':
			rank = 13
		case 'Q':
			rank = 12
		case 'J':
			rank = 11
		case 'T':
			rank = 10
		default:
			rank = int(card[0] - '0')
		}

		suit := ""
		switch card[1] {
		case 'H':
			suit = "Hearts"
		case 'D':
			suit = "Diamonds"
		case 'C':
			suit = "Clubs"
		case 'S':
			suit = "Spades"
		}

		result[i] = Card{Rank: rank, Suit: suit}
	}
	return result
}
