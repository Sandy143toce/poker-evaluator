package models

type PokerEvaluationResponse struct {
	Hand            string                `json:"hand"`
	HandRank        int                   `json:"handRank"`
	Cards           []string              `json:"cards"`
	PotentialBetter []PotentialBetterHand `json:"potentialBetterHands"`
}

type PotentialBetterHand struct {
	Name        string   `json:"name"`
	HandRank    int      `json:"handRank"`
	NeededCards []string `json:"neededCards"`
}
