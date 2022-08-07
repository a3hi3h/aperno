-- name: CreateExamRound :one
INSERT INTO examround (
  id,
  type,
  level,
  time,
  userid,
  orgid
) VALUES (
    $1, $2, $3, $4, $5, $6
) RETURNING *;

-- name: GetExamRound :one
SELECT * FROM examround
WHERE id = $1 LIMIT 1;
