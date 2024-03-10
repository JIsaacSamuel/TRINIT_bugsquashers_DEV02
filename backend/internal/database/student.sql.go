// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: student.sql

package database

import (
	"context"

	"github.com/google/uuid"
)

const createStudent = `-- name: CreateStudent :one
INSERT INTO student (id, name, emailID, passcode)
VALUES ($1, $2, $3, $4)
RETURNING id, name, emailid, passcode
`

type CreateStudentParams struct {
	ID       uuid.UUID
	Name     string
	Emailid  string
	Passcode string
}

func (q *Queries) CreateStudent(ctx context.Context, arg CreateStudentParams) (Student, error) {
	row := q.db.QueryRowContext(ctx, createStudent,
		arg.ID,
		arg.Name,
		arg.Emailid,
		arg.Passcode,
	)
	var i Student
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Emailid,
		&i.Passcode,
	)
	return i, err
}

const getStudentPass = `-- name: GetStudentPass :one
SELECT id, passcode from student where emailID = $1
`

type GetStudentPassRow struct {
	ID       uuid.UUID
	Passcode string
}

func (q *Queries) GetStudentPass(ctx context.Context, emailid string) (GetStudentPassRow, error) {
	row := q.db.QueryRowContext(ctx, getStudentPass, emailid)
	var i GetStudentPassRow
	err := row.Scan(&i.ID, &i.Passcode)
	return i, err
}
