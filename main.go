// 启动程序
package main

import (
	"context"
	"mis-server/common/config"
	"mis-server/pkg/db"
	"mis-server/pkg/log"
	"mis-server/pkg/redis"
	"mis-server/router"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	log := log.Log()
	gin.SetMode(config.Config.Server.Model)
	router := router.InitRouter()
	srv := &http.Server{
		Addr:    config.Config.Server.Address,
		Handler: router,
	}
	// 启动服务
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Infof("listen: %s\n", err)
		}
		log.Infof("listen %s\n", config.Config.Server.Address)
	}()
	quit := make(chan os.Signal, 1)
	// 监听消息
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
	log.Info("Shutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Info("Server exiting")
}

func init() {
	db.SetupDBLink()
	redis.SetupRedisDb()
}
