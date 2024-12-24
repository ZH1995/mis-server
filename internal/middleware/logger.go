package middleware

import (
	"MyGin/pkg/util"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Trace() gin.HandlerFunc {
	return func(c *gin.Context) {
		traceID := c.GetHeader("X-Trace-ID")
		if traceID == "" {
			traceID = uuid.New().String()
		}

		// 添加调试日志
		fmt.Printf("Generated trace_id: %s\n", traceID)

		c.Set(util.TraceIDKey, traceID)
		c.Header("X-Trace-ID", traceID)

		// 不再在这里创建 logger，让 util.GetContextLogger 来处理
		c.Next()
	}
}
