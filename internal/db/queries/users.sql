-- name: GetUserByID :one
SELECT *
FROM users
WHERE id = $1;

-- name: GetUserByAPIKey :one
SELECT *
FROM users
WHERE api_key = $1;

-- name: CreateUser :one
INSERT INTO users (email, api_key)
VALUES ($1, $2)
RETURNING *;