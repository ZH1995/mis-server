package util

import (
	"MyGin/global"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

const (
	TraceIDKey = "trace_id"
	LoggerKey  = "logger"
)

// GetContextLogger 从 gin.Context 获取带 trace_id 的 logger
func GetContextLogger(c *gin.Context) *zap.Logger {
	// 优先从 context 中获取已注入的 logger
	if logger, exists := c.Get(LoggerKey); exists {
		return logger.(*zap.Logger)
	}

	// 如果没有注入的 logger，创建一个新的
	logger := global.Logger
	if traceID, exists := c.Get(TraceIDKey); exists && traceID.(string) != "" {
		logger = logger.With(zap.String(TraceIDKey, traceID.(string)))
		// 缓存到 context 中
		c.Set(LoggerKey, logger)
		return logger
	}

	// 如果没有 traceID，直接返回全局 logger
	return logger
}

// L 是 GetContextLogger 的简写方法
func L(c *gin.Context) *zap.Logger {
	return GetContextLogger(c)
}
