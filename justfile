build:
    go build -o bin/dungeon-time-api cmd/main.go

run:
    go run cmd/main.go

migrate args:
    go run github.com/golang-migrate/migrate/v4/cmd/migrate@latest {{args}}
