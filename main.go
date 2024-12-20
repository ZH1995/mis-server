package main

import (
	"MyGin/config"
	"MyGin/router"
)

func main() {
	config.InitConfig()

	r := router.SetupRouter()

	// 启动HTTP服务器，默认监听在0.0.0.0:8080
	port := config.AppConfig.App.Port
	if port == "" {
		port = ":8080"
	}
	r.Run(port)
}
