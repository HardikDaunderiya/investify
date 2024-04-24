// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: user.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createUser = `-- name: CreateUser :one
INSERT INTO
    bk_users (
        user_email,
        user_password,
        user_phone_number,
        users_role_id,
        users_photo_link
    )
VALUES
    ($1, $2, $3, $4, $5) RETURNING user_id, user_email, user_phone_number, user_password, users_role_id, users_photo_link, created_at, updated_at, deleted_at
`

type CreateUserParams struct {
	UserEmail       string      `json:"user_email"`
	UserPassword    string      `json:"user_password"`
	UserPhoneNumber pgtype.Text `json:"user_phone_number"`
	UsersRoleID     int32       `json:"users_role_id"`
	UsersPhotoLink  pgtype.Text `json:"users_photo_link"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (BkUser, error) {
	row := q.db.QueryRow(ctx, createUser,
		arg.UserEmail,
		arg.UserPassword,
		arg.UserPhoneNumber,
		arg.UsersRoleID,
		arg.UsersPhotoLink,
	)
	var i BkUser
	err := row.Scan(
		&i.UserID,
		&i.UserEmail,
		&i.UserPhoneNumber,
		&i.UserPassword,
		&i.UsersRoleID,
		&i.UsersPhotoLink,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}