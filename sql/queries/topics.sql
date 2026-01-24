-- name: CreateTopic :one
INSERT INTO topics (id, name, description, created_at, updated_at)
VALUES (
    $1,
    $2,
    $3,
    NOW(),
    NOW()
)
RETURNING *;