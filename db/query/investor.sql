-- name: CreateInvestor :one
INSERT INTO
    bk_investor (
        investor_name,
        investor_user_id,
        investor_address_id
    )
VALUES
    ($1, $2, $3) RETURNING *;