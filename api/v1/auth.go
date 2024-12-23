package v1

import (
	"MyGin/global"
	"MyGin/internal/model"
	"MyGin/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Register(ctx *gin.Context) {
	var user model.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := service.Register(&user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Registration successful",
	})
}

func Login(ctx *gin.Context) {
	var input struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := ctx.ShouldBindJSON(&input); err != nil {
		global.Logger.Error("Login bind json error",
			zap.Error(err),
			zap.String("raw_data", ctx.Request.PostForm.Encode()),
		)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	token, err := service.Login(input.Username, input.Password)
	if err != nil {
		global.Logger.Error("Login failed",
			zap.String("username", input.Username),
			zap.Error(err),
		)
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	global.Logger.Info("Login success",
		zap.String("username", input.Username),
	)
	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
