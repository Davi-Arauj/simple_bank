-- name: CreateEntry :one
INSERT INTO entries (
  amount
) VALUES (
  $1
) RETURNING *;