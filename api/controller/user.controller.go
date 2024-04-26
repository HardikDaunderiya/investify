package controller

import (
	"investify/api/services"
	"investify/api/types"
	db "investify/db/sqlc"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// log.Print("i am in controllers here")

type UserController struct {
	store   db.Store
	userSrv services.UserService
}

func NewUserController(store db.Store, userSrv services.UserService) *UserController {
	return &UserController{store: store, userSrv: userSrv}
}

// var userService ser

func (u *UserController) Test(c *gin.Context) {
	log.Print("i am in controllers")
	c.JSON(http.StatusAccepted, gin.H{"message": "Everything ok"})
}

//logically
// 1. Parse the request
//initiate the transaction
//2. Extract the adress,user,Profile
//first create Adress
//create the user
//according to the role id in the user create the type of profile investor or owner
//commit the transaction

func (u *UserController) CreateUser(ctx *gin.Context) {

	var req types.CreateUserRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, types.GenerateErrorResponse(err, http.StatusBadRequest, "position 1"))
		return
	}

	respObject, err := u.userSrv.CreateUserService(ctx, req) // Delegate creation logic to user service
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, types.GenerateErrorResponse(err, http.StatusInternalServerError, "position 3"))
		return
	}

	ctx.JSON(http.StatusOK, types.GenerateResponse(respObject, http.StatusOK))

}

func (u *UserController) LoginUser(ctx *gin.Context) {
	var req types.LoginUserRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, types.GenerateErrorResponse(err, http.StatusBadRequest, "position 1"))
		return
	}

	respObject, err := u.userSrv.LoginUserService(ctx, req) // Delegate creation logic to user service
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, types.GenerateErrorResponse(err, http.StatusInternalServerError, "position 3"))
		return
	}

	ctx.JSON(http.StatusOK, types.GenerateResponse(respObject, http.StatusOK))
}
