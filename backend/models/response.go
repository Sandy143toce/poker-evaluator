package models

type PokerEvaluationResponse struct {
	Hand     string   `json:"hand"`
	HandRank int      `json:"handRank"`
	Cards    []string `json:"cards"`
}

// Add other response models here if needed
