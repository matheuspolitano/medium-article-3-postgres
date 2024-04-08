// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: user.sql

package db

import (
	"context"
	"database/sql"
)

const getUser = `-- name: GetUser :one
SELECT id, username, role, created_at FROM users
WHERE id = $1
`

func (q *Queries) GetUser(ctx context.Context, id int32) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Role,
		&i.CreatedAt,
	)
	return i, err
}

const insertUser = `-- name: InsertUser :one
INSERT INTO users(
  "username",
  "role"
)  VALUES(
    $1, $2
)
RETURNING id, username, role, created_at
`

type InsertUserParams struct {
	Username sql.NullString
	Role     sql.NullString
}

func (q *Queries) InsertUser(ctx context.Context, arg InsertUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, insertUser, arg.Username, arg.Role)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Role,
		&i.CreatedAt,
	)
	return i, err
}