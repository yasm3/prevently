-- name: CreatePush :one
INSERT INTO pushes (user_id, message)
VALUES ($1, $2)
RETURNING *;