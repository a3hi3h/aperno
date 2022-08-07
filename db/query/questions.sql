-- name: CreateQuestions :one
INSERT INTO questions (
  id,
  type,
  level,
  format,
  userid,
  question
) VALUES (
    $1, $2, $3, $4, $5, $6
) RETURNING *;

-- name: GetQuestion :one
SELECT * FROM questions where id = $1 LIMIT 1;
