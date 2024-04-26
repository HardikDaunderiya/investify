package middleware

import (
	"investify/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

// check for valid admin token
func JWTRestaurantAuth() gin.HandlerFunc {
	return func(context *gin.Context) {
		err := util.ValidateJWT(context)
		if err != nil {
			context.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication required"})
			context.Abort()
			return
		}
		error := util.ValidateOwnerRoleJWT(context)
		if error != nil {
			context.JSON(http.StatusUnauthorized, gin.H{"error": "Only Restaurunt is allowed to perform this action"})
			context.Abort()
			return
		}
		context.Next()
	}
}

// check for valid customer token
func JWTAuthCustomer() gin.HandlerFunc {
	return func(context *gin.Context) {
		err := util.ValidateJWT(context)
		if err != nil {
			context.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication required"})
			context.Abort()
			return
		}
		error := util.ValidateInvestorRoleJWT(context)
		if error != nil {
			context.JSON(http.StatusUnauthorized, gin.H{"error": "Only  Customers are allowed to perform this action"})
			context.Abort()
			return
		}
		context.Next()
	}
}
