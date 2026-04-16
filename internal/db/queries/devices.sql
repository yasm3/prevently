-- name: CreateDevice :one
INSERT INTO devices (user_id, name, type, config)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: ListDevicesByUser :many
SELECT *
FROM devices
WHERE user_id = $1;
