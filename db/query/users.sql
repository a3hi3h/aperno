-- name: CreateUser :one
insert into users (
   id, 
   status,
   orgid,
   first_name, 
   last_name,
   email, 
   hashedpwd,
   type
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8
) RETURNING *;

-- name: GetUserFromEmail :one
SELECT * FROM users WHERE email = $1 LIMIT 1;

