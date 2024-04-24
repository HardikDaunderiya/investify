-- name: CreateUser :one
INSERT INTO
    bk_users (
        user_email,
        user_password,
        user_phone_number,
        users_role_id,
        users_photo_link
    )
VALUES
    ($1, $2, $3, $4, $5) RETURNING *;