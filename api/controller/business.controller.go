package controller

import (
	"investify/api/services"
	db "investify/db/sqlc"
)

type BusinessController struct {
	store       db.Store
	businessSrv services.BusinessService
}

func NewBusinessController(store db.Store, BusinessSrv services.BusinessService) *BusinessController {
	return &BusinessController{store: store, businessSrv: BusinessSrv}
}
