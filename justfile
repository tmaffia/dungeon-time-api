default: lint run

init:
    go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

lint:
    go run honnef.co/go/tools/cmd/staticcheck@latest ./...

build:
    go build -o bin/dungeon-time-api cmd/main.go

run:
    go run cmd/main.go

test:
    go test ./...

mocks:
    go run github.com/vektra/mockery/v2@v2.44.1

migrate-create *args:
    migrate create -ext sql -dir db/migrations -seq {{args}}

migrate *args:
    migrate -database $DUNGEON_TIME_API_DATABASE_URL -path db/migrations {{args}} 

sqlc *args:
    go run github.com/sqlc-dev/sqlc/cmd/sqlc@latest {{args}}
