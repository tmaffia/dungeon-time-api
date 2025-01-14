-- name: GetUsers :many
SELECT id, username, email, roles, timezone FROM users;

-- name: GetUserByID :one
SELECT id, username, email, roles, timezone FROM users
WHERE id = $1 LIMIT 1;

-- name: GetUserByEmail :one
SELECT id, username, email, roles, timezone FROM users
WHERE email = $1 LIMIT 1;

-- name: GetUserByUsername :one
SELECT id, username, email, roles, timezone FROM users
WHERE username = $1 LIMIT 1;

-- name: GetUserFullByEmail :one
SELECT * FROM users
WHERE email = $1 LIMIT 1;

-- name: CreateUser :one
INSERT INTO users (username, email, password_hash, roles, timezone)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;
