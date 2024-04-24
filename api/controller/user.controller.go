package controller

import (
	"investify/api/types"
	db "investify/db/sqlc"
	"investify/util"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// log.Print("i am in controllers here")

type UserController struct {
	store db.Store
}

func NewUserController(store db.Store) *UserController {
	return &UserController{store: store}
}
func (u *UserController) Test(c *gin.Context) {
	log.Print("i am in controllers")
	c.JSON(http.StatusAccepted, gin.H{"message": "Everything ok"})
}

func (u *UserController) CreateUser(ctx *gin.Context) {
	var req types.CreateUserRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		// Handle parsing/binding error
		ctx.JSON(http.StatusBadRequest, types.GenerateErrorResponse(err, http.StatusBadRequest, "position 1"))
		return
	}
	hashedPassword, err := util.HashPassword(req.UserDetails.UserPassword)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, types.GenerateErrorResponse(err, http.StatusInternalServerError, "position 2"))
		return
	}

	arg := db.CreateUserParams{
		UserEmail:    req.UserDetails.UserEmail,
		UserPassword: hashedPassword,
		// UserPhoneNumber: req.UserDetails.UserPhoneNumber,
		// UsersPhotoLink: req.UserDetails.UsersPhotoLink,
		UsersRoleID: int32(req.UserDetails.UserRoleID),
	}
	user, err := u.store.CreateUser(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, types.GenerateErrorResponse(err, http.StatusInternalServerError, "position 3"))
		return
	}

	ctx.JSON(http.StatusOK, types.GenerateResponse(user, http.StatusOK))

}
