package api

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/tmaffia/dungeon-time-api/internal/service"
)

func StartApi() {
	conf := newConfig()
	dbpool, err := pgxpool.New(context.Background(), conf.databaseUrl)
	if err != nil {
		panic(err)
	}

	userService := service.NewUserService(dbpool)

	as := appState{
		userService: userService,
	}

	mux := http.NewServeMux()

	mux.HandleFunc("GET /api/v1/health", healthHandler)
	mux.HandleFunc("GET /api/v1/users", as.getUsersHandler)
	mux.HandleFunc("GET /api/v1/users/{id}", as.getUserHandler)

	log.Println("Starting Dungeon Time API on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func (as appState) getUsersHandler(w http.ResponseWriter, r *http.Request) {
	users, err := as.userService.GetUsers(context.Background())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	userJson, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(userJson)
}

func (as appState) getUserHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	user, err := as.userService.GetUserByID(context.Background(), int32(id))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	userJson, err := json.Marshal(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(userJson)
}
