package database

import (
	"context"
	"encoding/json"
	"time"

	"github.com/Sandy143toce/poker-evaluator/backend/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

func StoreGameResult(db *pgxpool.Pool, result models.PokerEvaluationResponse) error {
	query := `
		INSERT INTO game_results (player_best_hand, other_best_hands, created_at)
		VALUES ($1, $2, $3)
	`
	playerBestHandJSON, err := json.Marshal(result.PlayerBestHand)
	if err != nil {
		return err
	}

	otherBestHandsJSON, err := json.Marshal(result.OtherBestHands)
	if err != nil {
		return err
	}

	_, err = db.Exec(context.Background(), query, playerBestHandJSON, otherBestHandsJSON, time.Now())
	return err
}

func GetRecentGameResults(db *pgxpool.Pool, limit int) ([]models.PokerEvaluationResponse, error) {
	query := `
		SELECT player_best_hand, other_best_hands
		FROM game_results
		ORDER BY created_at DESC
		LIMIT $1
	`
	rows, err := db.Query(context.Background(), query, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []models.PokerEvaluationResponse
	for rows.Next() {
		var result models.PokerEvaluationResponse
		var playerBestHandJSON, otherBestHandsJSON []byte

		err := rows.Scan(&playerBestHandJSON, &otherBestHandsJSON)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(playerBestHandJSON, &result.PlayerBestHand)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(otherBestHandsJSON, &result.OtherBestHands)
		if err != nil {
			return nil, err
		}

		results = append(results, result)
	}

	return results, nil
}
