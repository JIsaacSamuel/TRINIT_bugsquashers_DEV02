-- name: CreateStudent :one
INSERT INTO student (id, name, emailID, passcode)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetStudentPass :one
SELECT id, passcode from student where emailID = $1;