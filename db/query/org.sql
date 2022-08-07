-- name: CreateOrg :one
INSERT INTO org (
  id,
  name,
  userid,
  accountid
) VALUES (
    $1, $2, $3, $4
) RETURNING *;

-- name: GetOrgFromUser :one
SELECT * FROM org
WHERE userid = $1 LIMIT 1;
