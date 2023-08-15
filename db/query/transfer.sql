-- Create
-- name: CreateTransfer :one
INSERT INTO transfers (
  from_account_id,
  to_account_id, 
  amount
) VALUES (
  $1, $2, $3
) RETURNING *;

-- Read One
-- name: GetTransfer :one
SELECT * FROM transfers
WHERE transfer_id = $1 LIMIT 1;

-- Read Many
-- name: ListTransfers :many
SELECT * FROM transfers
WHERE 
    from_account_id = $1 OR
    to_account_id = $2
ORDER BY transfer_id
LIMIT $3 
OFFSET $4;
