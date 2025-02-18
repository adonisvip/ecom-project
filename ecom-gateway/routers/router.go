package routers

import (
	"github.com/gin-gonic/gin"
	"ecom-gateway/handlers"
	"ecom-gateway/middleware"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.POST("/login", handlers.LoginHandler)
	protected := router.Group("/")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/product/:id", handlers.GetProductHandler)
	}

	return router
}
