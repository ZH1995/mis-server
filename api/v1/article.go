package v1

import (
	"MyGin/global"
	"MyGin/internal/model"
	"MyGin/internal/service"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateArticle(ctx *gin.Context) {
	var article model.Article
	if err := ctx.ShouldBindJSON(&article); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	article.CreateTime = time.Now().Unix()
	if err := service.CreateArticle(&article); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"article": article,
	})
}

func GetArticles(ctx *gin.Context) {
	articles, err := service.GetArticles()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			global.Logger.Error("111")
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		}
		return
	}
	ctx.JSON(
		http.StatusOK,
		gin.H{
			"articles": articles,
		},
	)
}

func GetArticleById(ctx *gin.Context) {
	id := ctx.Param("id")
	article, err := service.GetArticleByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		}
		return
	}
	ctx.JSON(
		http.StatusOK,
		gin.H{
			"article": article,
		},
	)
}
