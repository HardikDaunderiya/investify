package services

import db "investify/db/sqlc"

type InvestorService interface {
}

type InvestorServiceImpl struct {
	store db.Store
}

func NewInvestorService(store db.Store) *InvestorServiceImpl {
	return &InvestorServiceImpl{store: store}
}
