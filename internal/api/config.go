package api

import (
	"os"

	"github.com/tmaffia/dungeon-time-api/internal/service"
)

type appState struct {
	userService service.UserService
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
