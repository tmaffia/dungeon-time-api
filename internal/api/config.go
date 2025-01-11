package api

import (
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/tmaffia/dungeon-time-api/internal/db"
)

type appState struct {
	dbPool *pgxpool.Pool
	q      *db.Queries
}

type config struct {
	databaseUrl string
}

func newConfig() *config {
	dbUrl := os.Getenv("DUNGEON_TIME_API_DATABASE_URL")
	if dbUrl == "" {
		panic("DUNGEON_TIME_API_DATABASE_URL is required")
	}
	return &config{
		databaseUrl: dbUrl,
	}
}
