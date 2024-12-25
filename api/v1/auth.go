package v1

import (
	"MyGin/internal/model"
	"MyGin/internal/service"
	"MyGin/pkg/util"
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

// LoginRequest 登录请求参数
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// LoginResponse 登录响应
type LoginResponse struct {
	Token string `json:"token"`
}

// ErrorResponse 错误响应
type ErrorResponse struct {
	Error string `json:"error"`
}

// @Summary      User Login
// @Description  Authenticate user and return JWT token
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        input body LoginRequest true "Login Credentials"
// @Success      200  {object}  LoginResponse
// @Failure      400  {object}  ErrorResponse
// @Failure      401  {object}  ErrorResponse
// @Router       /auth/login [post]
func Login(ctx *gin.Context) {
	var input LoginRequest
	if err := ctx.ShouldBindJSON(&input); err != nil {
		util.L(ctx).Error("Login bind json error",
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
		util.L(ctx).Error("Login failed",
			zap.String("username", input.Username),
			zap.Error(err),
		)
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	util.L(ctx).Info("Login success",
		zap.String("username", input.Username),
	)
	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
