-- name: GetUsers :many
SELECT * FROM users;

-- name: GetUserByID :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: CreateUser :one
INSERT INTO users (username, email, password_hash, roles, timezone)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;
