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
) RETURNING id, status, type, orgid, first_name, last_name, email, hashedpwd, created_at, last_modified
`

type CreateUserParams struct {
	ID        uuid.UUID     `json:"id"`
	Status    int32         `json:"status"`
	Orgid     uuid.NullUUID `json:"orgid"`
	FirstName string        `json:"first_name"`
	LastName  string        `json:"last_name"`
	Email     string        `json:"email"`
	Hashedpwd string        `json:"hashedpwd"`
	Type      int32         `json:"type"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.ID,
		arg.Status,
		arg.Orgid,
		arg.FirstName,
		arg.LastName,
		arg.Email,
		arg.Hashedpwd,
		arg.Type,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Status,
		&i.Type,
		&i.Orgid,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.Hashedpwd,
		&i.CreatedAt,
		&i.LastModified,
	)
	return i, err
}

const getUserFromEmail = `-- name: GetUserFromEmail :one
SELECT id, status, type, orgid, first_name, last_name, email, hashedpwd, created_at, last_modified FROM users WHERE email = $1 LIMIT 1
`

func (q *Queries) GetUserFromEmail(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserFromEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Status,
		&i.Type,
		&i.Orgid,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.Hashedpwd,
		&i.CreatedAt,
		&i.LastModified,
	)
	return i, err
}
