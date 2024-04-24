-- name: Createaddress :one
INSERT INTO
    bk_address (
        address_street,
        address_city,
        address_state,
        address_country,
        address_zipcode 
    )
VALUES
    ($1, $2, $3, $4, $5) RETURNING *;