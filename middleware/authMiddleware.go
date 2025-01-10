// 鉴权中间件
package middleware

import (
	"mis-server/common/constant"
	"mis-server/common/result"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			result.Failed(c, int(result.APICode.NOAUTH), result.APICode.GetMessage(result.APICode.NOAUTH))
			c.Abort()
			return
		}
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			result.Failed(c, int(result.APICode.AUTHFORMATERROR), result.APICode.GetMessage(result.APICode.AUTHFORMATERROR))
			c.Abort()
			return
		}
		// todo 校验token
		var token = "token"
		c.Set(constant.ContextKeyUserObj, token)
		c.Next()
	}
}
