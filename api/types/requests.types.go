package types

type CreateUserRequest struct {
	ProfileDetails ProfileInfo
	UserDetails    UserType
	AdressDetails  AddressType
}

type LoginUserRequest struct {
	UserEmail    string `json:"user_email"`
	UserPassword string `json:"user_password"`
}
