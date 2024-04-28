package api

import (
	"investify/api/controller"
	"investify/api/middleware"
	"investify/api/services"
	"log"

	"github.com/gin-gonic/gin"
)

func SetupRouter(server *Server) *gin.Engine {
	router := server.router

	// Log statement for debugging

	// router.POST("/test", Test)

	//create a owner service
	//create a investor service
	//create a restaurant service

	v1 := router.Group("/api/v1")
	{
		// Log statement for debugging
		log.Println("Setting up API version 1 routes...")

		//Services
		userService := services.NewUserService(server.store)
		// ownerService := services.NewOwnerService(server.store)
		investorService := services.NewInvestorService(server.store)
		businessService := services.NewBusinessService(server.store)

		// Controllers
		userController := controller.NewUserController(server.store, userService)
		// ownerController := controller.NewOwnerController(server.store, ownerService)
		investorController := controller.NewInvestorController(server.store, investorService)
		businessController := controller.NewBusinessController(server.store, businessService)

		// Define routes for users
		users := v1.Group("/users")
		{
			// Log statement for debugging
			log.Println("Setting up user routes...")

			users.POST("/signup", userController.CreateUser)
			users.POST("/login", userController.LoginUser)
			users.GET("/test", userController.Test)
		}
		// owner := v1.Group("/owner")
		// {

		// }
		investor := v1.Group("/investor")
		{
			investor.Use(middleware.JWTAuthInvestor())
			investor.GET("/feed", investorController.GetBusinessFeedController)
			investor.GET("/:id", investorController.GetInvestorByIdController)

		}
		business := v1.Group("/business")
		{
			business.Use(middleware.JWTOwnerAuth())
			business.POST("/createBusiness", businessController.CreateBusiness)
			business.GET("/:id", businessController.GetBusinessByIdController)
			business.GET("/owner", businessController.GetBusinessByOwnerController)
			business.GET("/feed", businessController.GetInvestorFeedController)

		}

	}

	// Log statement for debugging
	log.Println("Router setup complete.")

	return router
}
