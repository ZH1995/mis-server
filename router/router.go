// 访问接口路由配置
package router

import (
	"mis-server/common/config"
	"mis-server/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 初始化路由
func InitRouter() *gin.Engine {
	router := gin.New()
	// 宕机恢复
	router.Use(gin.Recovery())
	router.Use(middleware.Cors())
	router.StaticFS(config.Config.ImageSettings.UploadDir, http.Dir(config.Config.ImageSettings.UploadDir))
	router.Use(gin.Logger())
	register(router)
	return router
}

// 路由注册
func register(router *gin.Engine) {
	// TODO 接口url写在此处
}
