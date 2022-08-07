-- name: CreateAccount :one
INSERT INTO accounts (
  id,
  name,
  status,
  type,
  userid,
  orgid
) VALUES (
    $1, $2, $3, $4, $5, $6
) RETURNING *;

-- name: GetAccount :one
SELECT * FROM accounts
WHERE userid = $1 LIMIT 1;
