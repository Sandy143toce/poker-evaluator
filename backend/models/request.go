package models

type PokerEvaluationRequest struct {
	PlayerCards []string `json:"playerCards"`
	TableCards  []string `json:"tableCards"`
}

// Add other request models here if needed
