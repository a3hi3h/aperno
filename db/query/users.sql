-- name: CreateUser :one
insert into users (
   u_uuid, 
   u_first_name, 
   u_last_name,
   u_type
) VALUES (
    $1, $2, $3, $4
) RETURNING *;


-- name: DeleteUser :exec
DELETE FROM users WHERE u_uuid = $1;

-- name: GetUser :one
SELECT * FROM users WHERE u_uuid = $1 LIMIT 1;