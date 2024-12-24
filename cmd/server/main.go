package main

import (
	v1 "MyGin/api/v1"
	config "MyGin/config"
	"MyGin/global"
	"MyGin/internal/middleware"
	"MyGin/pkg/util"
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	config.InitConfig()
	defer global.Logger.Sync()

	r := gin.New()
	r.Use(middleware.Trace())
	//r.Use(ginzap.Ginzap(global.Logger, time.RFC3339, true))
	// 修改日志中间件
	r.Use(func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery

		c.Next()

		logger := util.L(c)
		latency := time.Since(start)

		if len(c.Errors) > 0 {
			for _, e := range c.Errors.Errors() {
				logger.Error(e)
			}
		} else {
			logger.Info("request completed",
				zap.Int("status", c.Writer.Status()),
				zap.String("method", c.Request.Method),
				zap.String("path", path),
				zap.String("query", query),
				zap.String("ip", c.ClientIP()),
				zap.String("user-agent", c.Request.UserAgent()),
				zap.Duration("latency", latency),
			)
		}
	})
	r.Use(ginzap.RecoveryWithZap(global.Logger, true))

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
