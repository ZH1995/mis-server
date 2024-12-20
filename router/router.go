package router

import (
	"MyGin/controller"
	"MyGin/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	auth := r.Group("/api/auth")
	{
		auth.POST("/login", controller.Login)
		auth.POST("/register", controller.Register)
	}

	api := r.Group("/api")
	api.GET("/exchange_rate", controller.GetExchangeRate)
	api.Use(middleware.AuthMiddleware())
	{
		api.POST("/exchange_rate", controller.CreateEchangeRate)
	}
	return r
}
