-- name: CreateUserSession :one
INSERT INTO usersessions (
  id,
  userid,
  refresh_token,
  user_agent,
  client_ip,
  expires_at,
  is_blocked
) VALUES (
  $1, $2, $3, $4, $5, $6, $7
) RETURNING *;

-- name: GetSession :one
SELECT * FROM usersessions
WHERE id = $1 LIMIT 1;

-- name: GetUserIdForSession :one
SELECT userid FROM usersessions
WHERE id = $1 LIMIT 1;

