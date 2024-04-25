package types

import "github.com/jackc/pgx/v5/pgtype"

type AddressType struct {
	AddressStreet  string      `json:"address_street" binding:"required"`
	AddressCity    string      `json:"address_city" binding:"required"`
	AddressState   string      `json:"address_state" binding:"required"`
	AddressCountry pgtype.Text `json:"address_country" binding:"required"`
	AddressZipcode string      `json:"address_zipcode" binding:"required"`
}

type UserType struct {
	UserEmail       string `json:"user_email" binding:"email,required"`
	UserPassword    string `json:"user_password" binding:"required"`
	UserRoleID      int32  `json:"user_role_id" binding:"required"`
	UserPhoneNumber string `json:"user_phone_number"`
	UsersPhotoLink  string `json:"users_photo_link"`
}
type ProfileInfo struct {
	ProfileName pgtype.Text `json:"profile_name" binding:"required"`
}
