package utils

import (
	"context"
	"log"

	"github.com/Sandy143toce/poker-evaluator/backend/database"
	"github.com/jackc/pgx/v5/pgxpool"
)

func InitDatabase() *pgxpool.Pool {
	db, err := database.InitDB()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// Test the connection
	err = db.Ping(context.Background())
	if err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	log.Println("Successfully connected to the database")
	return db
}
