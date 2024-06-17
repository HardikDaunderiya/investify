// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package db

import (
	"context"
)

type Querier interface {
	CreateAddress(ctx context.Context, arg CreateAddressParams) (BkAddress, error)
	CreateBusiness(ctx context.Context, arg CreateBusinessParams) (BkBusiness, error)
	CreateInvestor(ctx context.Context, arg CreateInvestorParams) (BkInvestor, error)
	CreateOwner(ctx context.Context, arg CreateOwnerParams) (BkOwner, error)
	CreateToken(ctx context.Context, arg CreateTokenParams) (BkToken, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (BkUser, error)
	GetAddressById(ctx context.Context, addressID int64) (BkAddress, error)
	GetBusinessById(ctx context.Context, businessID int64) (BkBusiness, error)
	GetBusinessByOwnerId(ctx context.Context, businessOwnerID int64) ([]BkBusiness, error)
	GetBusinessFeed(ctx context.Context) ([]BkBusiness, error)
	GetInvestorById(ctx context.Context, investorID int64) (BkInvestor, error)
	GetInvestorByUserId(ctx context.Context, investorUserID int64) (BkInvestor, error)
	GetInvestorFeed(ctx context.Context) ([]BkInvestor, error)
	GetOwnerByUserId(ctx context.Context, ownerUserID int64) (BkOwner, error)
	GetUserByEmail(ctx context.Context, userEmail string) (BkUser, error)
	GetUserById(ctx context.Context, userID int64) (BkUser, error)
}

var _ Querier = (*Queries)(nil)
