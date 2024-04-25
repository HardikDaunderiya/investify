package services

import (
	db "investify/db/sqlc"
	"investify/util"

	"investify/api/types"

	"github.com/gin-gonic/gin"
)

type UserService interface {
	CreateUser(ctx *gin.Context, req types.CreateUserRequest) (types.CreateUserResponse, error)
}

type UserServiceImpl struct {
	store db.Store
}

func NewUserService(store db.Store) *UserServiceImpl {
	return &UserServiceImpl{store: store}
}

func (u *UserServiceImpl) CreateUserService(ctx *gin.Context, req types.CreateUserRequest) (types.CreateUserResponse, error) {
	// Implement the logic here
	hashedPassword, err := util.HashPassword(req.UserDetails.UserPassword)
	if err != nil {
		return types.CreateUserResponse{}, err
	}

	var respObject types.CreateUserResponse

	err = u.store.ExecTx(ctx, func(tx *db.Queries) error {
		address, err := u.store.CreateAddress(ctx, db.CreateAddressParams{
			AddressStreet:  req.AdressDetails.AddressStreet,
			AddressCity:    req.AdressDetails.AddressCity,
			AddressState:   req.AdressDetails.AddressState,
			AddressCountry: req.AdressDetails.AddressCountry,
			AddressZipcode: req.AdressDetails.AddressZipcode,
		})
		if err != nil {
			return err
		}
		respObject.AddressInfo = address

		user, err := u.store.CreateUser(ctx, db.CreateUserParams{
			UserEmail:    req.UserDetails.UserEmail,
			UserPassword: hashedPassword,
			UsersRoleID:  req.UserDetails.UserRoleID,
		})
		if err != nil {
			return err
		}
		respObject.UserInfo = user

		if user.UsersRoleID == 1 {
			owner, err := u.store.CreateOwner(ctx, db.CreateOwnerParams{
				OwnerName:      req.ProfileDetails.ProfileName,
				OwnerAddressID: address.AddressID,
				OwnerUserID:    user.UserID,
			})
			if err != nil {
				return err
			}
			respObject.ProfileInfo = owner

		} else if user.UsersRoleID == 2 {
			investor, err := u.store.CreateInvestor(ctx, db.CreateInvestorParams{
				InvestorName:      req.ProfileDetails.ProfileName,
				InvestorAddressID: address.AddressID,
				InvestorUserID:    user.UserID,
			})
			if err != nil {
				return err
			}
			respObject.ProfileInfo = investor
		}

		return nil // Commit transaction
	})

	if err != nil {
		return types.CreateUserResponse{}, err
	}

	// Return success response
	return respObject, nil

}
