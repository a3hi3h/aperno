// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0
// source: users.sql

package db

import (
	"context"

	"github.com/google/uuid"
)

const createUser = `-- name: CreateUser :one
insert into users (
   u_uuid, 
   u_first_name, 
   u_last_name,
   u_type
) VALUES (
    $1, $2, $3, $4
) RETURNING u_uuid, u_first_name, u_last_name, u_type, u_org, u_detail
`

type CreateUserParams struct {
	UUuid      uuid.UUID `json:"u_uuid"`
	UFirstName string    `json:"u_first_name"`
	ULastName  string    `json:"u_last_name"`
	UType      int32     `json:"u_type"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.UUuid,
		arg.UFirstName,
		arg.ULastName,
		arg.UType,
	)
	var i User
	err := row.Scan(
		&i.UUuid,
		&i.UFirstName,
		&i.ULastName,
		&i.UType,
		&i.UOrg,
		&i.UDetail,
	)
	return i, err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM users WHERE u_uuid = $1
`

func (q *Queries) DeleteUser(ctx context.Context, uUuid uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteUser, uUuid)
	return err
}

const getUser = `-- name: GetUser :one
SELECT u_uuid, u_first_name, u_last_name, u_type, u_org, u_detail FROM users WHERE u_uuid = $1 LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, uUuid uuid.UUID) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, uUuid)
	var i User
	err := row.Scan(
		&i.UUuid,
		&i.UFirstName,
		&i.ULastName,
		&i.UType,
		&i.UOrg,
		&i.UDetail,
	)
	return i, err
}
