package services

import (
	db "investify/db/sqlc"
	"investify/util"

	"investify/api/types"

	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type UserService interface {
	CreateUserService(ctx *gin.Context, req types.CreateUserRequest) (types.CreateUserResponse, error)
	LoginUserService(ctx *gin.Context, req types.LoginUserRequest) (types.LoginUserResponse, error)
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

func (u *UserServiceImpl) LoginUserService(ctx *gin.Context, req types.LoginUserRequest) (types.LoginUserResponse, error) {
	// Implement the logic here
	//check user
	//verify password

	var LoginUserResponse types.LoginUserResponse

	user, err := u.store.GetUserByEmail(ctx, req.UserEmail)
	if err != nil {
		return types.LoginUserResponse{}, err
	}

	err = util.CheckPassword(req.UserPassword, user.UserPassword)
	if err != nil {
		return types.LoginUserResponse{}, err
	}

	token, err := util.GenerateJWT(user)
	if err != nil {
		return types.LoginUserResponse{}, err

	}
	LoginUserResponse.AccessToken = token
	uuidToken := uuid.New()
	refreshToken, err := u.store.CreateToken(ctx, db.CreateTokenParams{
		TokenValue:      uuidToken,
		TokenUserID:     user.UserID,
		TokenExpiryDate: pgtype.Timestamptz{Time: time.Now().Add(7 * 24 * time.Hour)},
	})
	LoginUserResponse.RefreshToken = refreshToken.TokenValue.String()
	if err != nil {
		return types.LoginUserResponse{}, err
	}
	return LoginUserResponse, nil
}
