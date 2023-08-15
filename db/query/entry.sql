-- Create
-- name: CreateEntrie :one
INSERT INTO entries (
  account_id,
  amount
) VALUES (
  $1, $2
) RETURNING *;

-- Read One
-- name: GetEntry :one
SELECT * FROM entries
WHERE entrie_id = $1 LIMIT 1;

-- Read Many
-- name: ListEntries :many
SELECT * FROM entries
WHERE account_id = $1
ORDER BY entrie_id
LIMIT $2 
OFFSET $3;
