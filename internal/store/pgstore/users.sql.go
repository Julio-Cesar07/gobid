// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: users.sql

package pgstore

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users ("username", "email", "password_hash", "bio")
VALUES ($1, $2, $3, $4)
RETURNING id
`

type CreateUserParams struct {
	Username     string      `json:"username"`
	Email        string      `json:"email"`
	PasswordHash []byte      `json:"password_hash"`
	Bio          pgtype.Text `json:"bio"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (uuid.UUID, error) {
	row := q.db.QueryRow(ctx, createUser,
		arg.Username,
		arg.Email,
		arg.PasswordHash,
		arg.Bio,
	)
	var id uuid.UUID
	err := row.Scan(&id)
	return id, err
}

const getUserById = `-- name: GetUserById :one

SELECT id, username, email, password_hash, bio, created_at, updated_at FROM users WHERE id = $1 LIMIT 1
`

func (q *Queries) GetUserById(ctx context.Context, id uuid.UUID) (User, error) {
	row := q.db.QueryRow(ctx, getUserById, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.PasswordHash,
		&i.Bio,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
