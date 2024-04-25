package api

import (
	"investify/api/controller"
	"investify/api/services"
	"log"

	"github.com/gin-gonic/gin"
)

func SetupRouter(server *Server) *gin.Engine {
	router := server.router

	// Log statement for debugging

	// router.POST("/test", Test)

	v1 := router.Group("/api/v1")
	{
		// Log statement for debugging
		log.Println("Setting up API version 1 routes...")

		//Services
		userService := services.NewUserService(server.store)
		// Controllers
		userController := controller.NewUserController(server.store, userService)

		// Define routes for users
		users := v1.Group("/users")
		{
			// Log statement for debugging
			log.Println("Setting up user routes...")

			users.POST("/signup", userController.CreateUser)
			users.GET("/test", userController.Test)
		}
	}

	// Log statement for debugging
	log.Println("Router setup complete.")

	return router
}
