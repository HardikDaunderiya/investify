-- name: CreateUser :one
INSERT INTO
    user_email,
VALUES
    ($ 1) RETURNING *;