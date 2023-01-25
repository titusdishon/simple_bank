// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: user.sql

package db

import (
	"context"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (
  username,
  full_name,
  email,
  phone_number,
  hashed_password
) VALUES (
  $1, $2, $3,$4,$5
)
RETURNING username, hashed_password, full_name, phone_number, email, password_changed_at, created_at
`

type CreateUserParams struct {
	Username       string `json:"username"`
	FullName       string `json:"full_name"`
	Email          string `json:"email"`
	PhoneNumber    string `json:"phone_number"`
	HashedPassword string `json:"hashed_password"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.Username,
		arg.FullName,
		arg.Email,
		arg.PhoneNumber,
		arg.HashedPassword,
	)
	var i User
	err := row.Scan(
		&i.Username,
		&i.HashedPassword,
		&i.FullName,
		&i.PhoneNumber,
		&i.Email,
		&i.PasswordChangedAt,
		&i.CreatedAt,
	)
	return i, err
}

const getUser = `-- name: GetUser :one
SELECT username, hashed_password, full_name, phone_number, email, password_changed_at, created_at FROM users
WHERE username = $1 LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, username string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, username)
	var i User
	err := row.Scan(
		&i.Username,
		&i.HashedPassword,
		&i.FullName,
		&i.PhoneNumber,
		&i.Email,
		&i.PasswordChangedAt,
		&i.CreatedAt,
	)
	return i, err
}
