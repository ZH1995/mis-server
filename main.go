package main

import (
	"MyGin/config"
	"fmt"
)

func main() {
	config.InitConfig()
	fmt.Println(config.AppConfig.App.Name)
	/*
		// 创建一个默认的Gin路由器
		r := gin.Default()

		r.Use(gin.Logger())

		r.Use(gin.Recovery())

		// 定义一个GET请求的路由
		r.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})

		r.POST("/post", func(c *gin.Context) {
			ids := c.QueryMap("ids")
			names := c.PostFormMap("names")
			fmt.Printf("ids: %v; names: %v\n", ids, names)
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

		r.GET("/welcome", func(c *gin.Context) {
			firstname := c.DefaultQuery("firstname", "Guest")
			lastname := c.Query("lastname")
			c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
		})

		r.POST("/form_post", func(c *gin.Context) {
			message := c.PostForm("message")
			nick := c.DefaultPostForm("nick", "anonymous")
			c.JSON(http.StatusOK, gin.H{
				"status":  "posted",
				"message": message,
				"nick":    nick,
			})
		})

		// 启动HTTP服务器，默认监听在0.0.0.0:8080
		r.Run()
	*/
}
