-- name: CreateTutor :one
INSERT INTO tutor (id, name, emailID, created_at, passcode)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetTutorPass :one
SELECT id, passcode from tutor where emailID = $1;

-- name: GetTutorCred :one
SELECT * from tutor where id = $1;