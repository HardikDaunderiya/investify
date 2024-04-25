package types

type CreateUserRequest struct {
	ProfileDetails ProfileInfo
	UserDetails    UserType
	AdressDetails  AddressType
}
