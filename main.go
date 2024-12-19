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

	r.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action
		c.String(http.StatusOK, message)
	})

	r.POST("/user/:name/*action", func(c *gin.Context) {
		b := c.FullPath() == "/user/:name/*action"
		c.String(http.StatusOK, "%t", b)
	})

	r.GET("/user/groups", func(c *gin.Context) {
		c.String(http.StatusOK, "The avaliable groups are [...]")
	})

	// 启动HTTP服务器，默认监听在0.0.0.0:8080
	r.Run()
}
