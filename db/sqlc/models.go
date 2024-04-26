// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package db

import (
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type BkAddress struct {
	AddressID      int64       `json:"address_id"`
	AddressStreet  string      `json:"address_street"`
	AddressCity    string      `json:"address_city"`
	AddressState   string      `json:"address_state"`
	AddressCountry pgtype.Text `json:"address_country"`
	AddressZipcode string      `json:"address_zipcode"`
}

type BkBusiness struct {
	BusinessID             int64            `json:"business_id"`
	BusinessOwnerID        int64            `json:"business_owner_id"`
	BusinessOwnerFirstname string           `json:"business_owner_firstname"`
	BusinessOwnerLastname  string           `json:"business_owner_lastname"`
	BusinessEmail          string           `json:"business_email"`
	BusinessContact        string           `json:"business_contact"`
	BusinessName           string           `json:"business_name"`
	BusinessAddressID      pgtype.Int8      `json:"business_address_id"`
	BusinessUserID         pgtype.Int8      `json:"business_user_id"`
	BusinessRatings        pgtype.Numeric   `json:"business_ratings"`
	BusinessMinAmount      pgtype.Numeric   `json:"business_minAmount"`
	CreatedAt              pgtype.Timestamp `json:"created_at"`
	UpdatedAt              pgtype.Timestamp `json:"updated_at"`
	DeletedAt              pgtype.Timestamp `json:"deleted_at"`
}

type BkInvestor struct {
	InvestorID        int64            `json:"investor_id"`
	InvestorName      pgtype.Text      `json:"investor_name"`
	InvestorUserID    int64            `json:"investor_user_id"`
	InvestorAddressID int64            `json:"investor_address_id"`
	CreatedAt         pgtype.Timestamp `json:"created_at"`
	UpdatedAt         pgtype.Timestamp `json:"updated_at"`
	DeletedAt         pgtype.Timestamp `json:"deleted_at"`
}

type BkOwner struct {
	OwnerID        int64            `json:"owner_id"`
	OwnerName      pgtype.Text      `json:"owner_name"`
	OwnerUserID    int64            `json:"owner_user_id"`
	OwnerAddressID int64            `json:"owner_address_id"`
	CreatedAt      pgtype.Timestamp `json:"created_at"`
	UpdatedAt      pgtype.Timestamp `json:"updated_at"`
	DeletedAt      pgtype.Timestamp `json:"deleted_at"`
}

type BkRole struct {
	RoleID   int32  `json:"role_id"`
	RoleName string `json:"role_name"`
}

type BkToken struct {
	TokenID         int64              `json:"token_id"`
	TokenValue      uuid.UUID          `json:"token_value"`
	TokenUserID     int64              `json:"token_user_id"`
	TokenExpiryDate pgtype.Timestamptz `json:"token_expiry_date"`
}

type BkUser struct {
	UserID          int64            `json:"user_id"`
	UserEmail       string           `json:"user_email"`
	UserPhoneNumber pgtype.Text      `json:"user_phone_number"`
	UserPassword    string           `json:"user_password"`
	UsersRoleID     int32            `json:"users_role_id"`
	UsersPhotoLink  pgtype.Text      `json:"users_photo_link"`
	CreatedAt       pgtype.Timestamp `json:"created_at"`
	UpdatedAt       pgtype.Timestamp `json:"updated_at"`
	DeletedAt       pgtype.Timestamp `json:"deleted_at"`
}
