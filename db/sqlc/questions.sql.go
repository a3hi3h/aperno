// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0
// source: questions.sql

package db

import (
	"context"

	"github.com/google/uuid"
)

const createQuestions = `-- name: CreateQuestions :one
INSERT INTO questions (
  id,
  type,
  level,
  format,
  userid,
  question
) VALUES (
    $1, $2, $3, $4, $5, $6
) RETURNING id, type, level, format, userid, orgid, question, question_time, created_at, last_modified
`

type CreateQuestionsParams struct {
	ID       uuid.UUID `json:"id"`
	Type     int32     `json:"type"`
	Level    int32     `json:"level"`
	Format   int32     `json:"format"`
	Userid   uuid.UUID `json:"userid"`
	Question string    `json:"question"`
}

func (q *Queries) CreateQuestions(ctx context.Context, arg CreateQuestionsParams) (Question, error) {
	row := q.db.QueryRowContext(ctx, createQuestions,
		arg.ID,
		arg.Type,
		arg.Level,
		arg.Format,
		arg.Userid,
		arg.Question,
	)
	var i Question
	err := row.Scan(
		&i.ID,
		&i.Type,
		&i.Level,
		&i.Format,
		&i.Userid,
		&i.Orgid,
		&i.Question,
		&i.QuestionTime,
		&i.CreatedAt,
		&i.LastModified,
	)
	return i, err
}

const getQuestion = `-- name: GetQuestion :one
SELECT id, type, level, format, userid, orgid, question, question_time, created_at, last_modified FROM questions where id = $1 LIMIT 1
`

func (q *Queries) GetQuestion(ctx context.Context, id uuid.UUID) (Question, error) {
	row := q.db.QueryRowContext(ctx, getQuestion, id)
	var i Question
	err := row.Scan(
		&i.ID,
		&i.Type,
		&i.Level,
		&i.Format,
		&i.Userid,
		&i.Orgid,
		&i.Question,
		&i.QuestionTime,
		&i.CreatedAt,
		&i.LastModified,
	)
	return i, err
}
