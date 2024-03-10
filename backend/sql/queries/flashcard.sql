-- name: ViewFlashCard :many
SELECT word, meaning from flashcard where course_id = $1;

-- name: CreateFlashCard :one
INSERT INTO flashcard (id, course_id, word, meaning)
VALUES ($1, $2, $3, $4)
RETURNING *;