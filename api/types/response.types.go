package types

import db "investify/db/sqlc"

type BaseHttpResponse struct {
	Status  string      `json:"status"`
	Data    interface{} `json:"data"`
	Message string      `json:message`
}
type BaseErrorResponse struct {
	Status  string      `json:"status"`
	Message interface{} `json:"error"`
	Extra   string      `json:"extra"`
}

func GenerateResponse(data interface{}, message string) *BaseHttpResponse {
	return &BaseHttpResponse{
		Status:  "success",
		Data:    data,
		Message: message,
	}
}

func GenerateErrorResponse(err error, statusCode int, extra string) *BaseErrorResponse {
	if extra == "" {
		extra = "" // Set extra to an empty string if it's not provided
	}
	return &BaseErrorResponse{
		Status:  "error",
		Message: err.Error(),
		Extra:   extra,
	}
}

type CreateUserResponse struct {
	UserInfo    db.BkUser    `json:"user_info"`
	AddressInfo db.BkAddress `json:"address_info"`
	ProfileInfo interface{}  `json:"profile_info"`
}
type LoginUserResponse struct {
	AccessToken  string `json:"access_token`
	RefreshToken string `json:"refresh_token`
}

type CreateBusinessResponse struct {
	BusinessInfo db.BkBusiness `json:"business_info"`
	AddressInfo  db.BkAddress  `json:"address_info"`
}

type GetBusinessResponse struct {
	BusinessInfo db.BkBusiness `json:"business_info"`
}
type GetBusinessFeedResponse struct {
	BusinessInfo []db.BkBusiness `json:"business_info"`
}
type GetInvestorResponse struct {
	InvestorInfo db.BkInvestor `json:"investor_info"`
}
type GetInvestorFeedResponse struct {
	InvestorInfo []db.BkInvestor `json:"investor_info"`
}
