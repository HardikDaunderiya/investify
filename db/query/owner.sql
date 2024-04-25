-- name: CreateOwner :one
INSERT INTO
    bk_owner (
        owner_name,
        owner_user_id,
        owner_address_id
    )
VALUES
    ($1, $2, $3) RETURNING *;