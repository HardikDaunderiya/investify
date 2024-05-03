package controller

import (
	"investify/api/services"
	"investify/api/types"
	db "investify/db/sqlc"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BusinessController struct {
	store       db.Store
	businessSrv services.BusinessService
}

func NewBusinessController(store db.Store, BusinessSrv services.BusinessService) *BusinessController {
	return &BusinessController{store: store, businessSrv: BusinessSrv}
}

func (b *BusinessController) CreateBusiness(ctx *gin.Context) {
	var req types.CreateBusinessRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, types.GenerateErrorResponse(err, http.StatusBadRequest, "position 1"))
		return
	}
	// fmt.Println(reflect.TypeOf(req.AdressDetails))
	_, err = b.businessSrv.CreateBusinessService(ctx, req)
	// Delegate creation logic to user service
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, types.GenerateErrorResponse(err, http.StatusInternalServerError, "position 3"))
		return
	}

	ctx.JSON(http.StatusOK, types.GenerateResponse(nil, "Business Created Sucessfully"))

}
func (b *BusinessController) GetBusinessByIdController(ctx *gin.Context) {

	respObject, err := b.businessSrv.GetBusinessService(ctx)
	// Delegate creation logic to user service
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, types.GenerateErrorResponse(err, http.StatusInternalServerError, "position 3"))
		return
	}

	ctx.JSON(http.StatusOK, types.GenerateResponse(respObject, "Business fetched"))

}

func (b *BusinessController) GetBusinessByOwnerController(ctx *gin.Context) {

	respObject, err := b.businessSrv.GetBusinessServiceByOwner(ctx)
	// Delegate cre	ation logic to user service
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, types.GenerateErrorResponse(err, http.StatusInternalServerError, "position 3"))
		return
	}

	ctx.JSON(http.StatusOK, types.GenerateResponse(respObject, "Business fetched"))

}
func (b *BusinessController) GetInvestorFeedController(ctx *gin.Context) {

	respObject, err := b.businessSrv.GetInvestorFeedService(ctx)
	// Delegate creation logic to user service
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, types.GenerateErrorResponse(err, http.StatusInternalServerError, "position 3"))
		return
	}

	ctx.JSON(http.StatusOK, types.GenerateResponse(respObject, "Business fetched"))

}
