package models

type Hand struct {
	Name     string   `json:"name"`
	Rank     int      `json:"rank"`
	Cards    []string `json:"cards"`
	Sequence []int    `json:"sequence"`
}

type PokerEvaluationResponse struct {
	PlayerBestHand Hand   `json:"playerBestHand"`
	OtherBestHands []Hand `json:"otherBestHands"`
}

// Add other response models here if needed
