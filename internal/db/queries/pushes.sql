-- name: CreatePush :one
INSERT INTO pushes (user_id, message)
VALUES ($1, $2)
RETURNING *;

-- name: ClaimPendingPushes :many
UPDATE pushes
SET status = 'processing'
WHERE id IN (
    SELECT id
    FROM pushes
    WHERE status = 'pending'
    ORDER BY created_at
    LIMIT $1
)
RETURNING *;

-- name: MarkPushFailed :one
UPDATE pushes
SET status = 'failed', last_error = $2, attempts = attempts + 1
WHERE id = $1
RETURNING *;

-- name: MarkPushSent :one
UPDATE pushes
SET status = 'sent', attempts = attempts + 1
WHERE id = $1
RETURNING *;