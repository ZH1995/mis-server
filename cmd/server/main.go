package main

import (
	v1 "MyGin/api/v1"
	config "MyGin/config"
	"MyGin/internal/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	config.InitConfig()
	r := gin.Default()

	auth := r.Group("/api/auth")
	{
		auth.POST("/login", v1.Login)
		auth.POST("/register", v1.Register)
	}

	api := r.Group("/api")
	api.GET("/exchange_rate", v1.GetExchangeRate)
	api.Use(middleware.AuthMiddleware())
	{
		api.POST("/exchange_rate", v1.CreateExchangeRate)
		api.POST("/articles", v1.CreateArticle)
		api.GET("/articles", v1.GetArticles)
		api.GET("/articles/:id", v1.GetArticleById)
		api.POST("/articles/:id/like", v1.LikeArticle)
		api.GET("/articles/:id/like", v1.GetArticleLikes)
	}
	r.Run(config.AppConfig.App.Port)
}
