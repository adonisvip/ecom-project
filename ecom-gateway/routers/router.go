package routers

import (
	"ecom-gateway/handlers"
	"ecom-gateway/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// API version 1 routes
	apiV1 := router.Group("/api/v1")
	{
		// Public routes
		apiV1.POST("/login", handlers.LoginHandler)
		apiV1.POST("/signup", handlers.SignupHandler)

		// Protected routes
		protected := apiV1.Group("/")
		protected.Use(middleware.AuthMiddleware())
		{
			protected.GET("/product/:id", handlers.GetProductHandler)
		}
	}

	return router
}
