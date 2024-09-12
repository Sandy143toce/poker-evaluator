package utils

import (
	"context"
	"encoding/json"
	"time"

	"github.com/Sandy143toce/poker-evaluator/backend/models"
	"github.com/go-redis/redis/v8"
)

const (
	RecentGameResultsKey = "recent_game_results"
	CacheExpiration      = 5 * time.Minute
)

func CacheRecentGameResults(client *redis.Client, results []models.PokerEvaluationResponse) error {
	data, err := json.Marshal(results)
	if err != nil {
		return err
	}

	return client.Set(context.Background(), RecentGameResultsKey, data, CacheExpiration).Err()
}

func GetCachedRecentGameResults(client *redis.Client) ([]models.PokerEvaluationResponse, error) {
	data, err := client.Get(context.Background(), RecentGameResultsKey).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, nil
		}
		return nil, err
	}

	var results []models.PokerEvaluationResponse
	err = json.Unmarshal([]byte(data), &results)
	if err != nil {
		return nil, err
	}

	return results, nil
}
