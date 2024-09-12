package database

import (
	"context"
	"time"

	"github.com/Sandy143toce/poker-evaluator/backend/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

func StoreGameResult(db *pgxpool.Pool, result models.PokerEvaluationResponse) error {
	query := `
		INSERT INTO game_results (hand, hand_rank, cards, created_at)
		VALUES ($1, $2, $3, $4)
	`
	_, err := db.Exec(context.Background(), query, result.Hand, result.HandRank, result.Cards, time.Now())
	return err
}

func GetRecentGameResults(db *pgxpool.Pool, limit int) ([]models.PokerEvaluationResponse, error) {
	query := `
		SELECT hand, hand_rank, cards
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
		err := rows.Scan(&result.Hand, &result.HandRank, &result.Cards)
		if err != nil {
			return nil, err
		}
		results = append(results, result)
	}

	return results, nil
}
