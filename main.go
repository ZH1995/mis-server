package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// 创建一个默认的Gin路由器
	r := gin.Default()

	// 定义一个GET请求的路由
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.POST("/post", func(c *gin.Context) {
		c.String(http.StatusOK, "This is a POST request")
	})

	// 启动HTTP服务器，默认监听在0.0.0.0:8080
	r.Run()
}
