// 通用返回结构
package result

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Result struct {
	Code    uint        `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Success(c *gin.Context, data interface{}) {
	if data == nil {
		data = gin.H{}
	}
	c.JSON(http.StatusOK, Result{
		Code:    APICode.SUCCESS,
		Message: APICode.GetMessage(APICode.SUCCESS),
		Data:    data,
	})
}

func Failed(c *gin.Context, code int, message string) {
	c.JSON(http.StatusOK, Result{
		Code:    uint(code),
		Message: message,
		Data:    gin.H{},
	})
}
