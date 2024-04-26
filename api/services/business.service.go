package services

import db "investify/db/sqlc"

type BusinessService interface {
}

type BusinessServiceImpl struct {
	store db.Store
}

func NewBusinessService(store db.Store) *BusinessServiceImpl {
	return &BusinessServiceImpl{store: store}
}
