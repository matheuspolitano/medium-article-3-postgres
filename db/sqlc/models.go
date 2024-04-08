// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package db

import (
	"database/sql"
	"time"
)

type Follow struct {
	FollowingUserID sql.NullInt32
	FollowedUserID  sql.NullInt32
	CreatedAt       time.Time
}

type Post struct {
	ID    int32
	Title sql.NullString
	// Content of the post
	Body      sql.NullString
	UserID    sql.NullInt32
	Status    sql.NullString
	CreatedAt time.Time
}

type User struct {
	ID        int32
	Username  sql.NullString
	Role      sql.NullString
	CreatedAt time.Time
}