package controller

import (
	"investify/api/services"
	db "investify/db/sqlc"
)

type InvestorController struct {
	store       db.Store
	investorSrv services.InvestorService
}

func NewInvestorController(store db.Store, InvestorSrv services.InvestorService) *InvestorController {
	return &InvestorController{store: store, investorSrv: InvestorSrv}
}
