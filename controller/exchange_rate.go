package controller

import (
	"MyGin/global"
	"MyGin/model"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateEchangeRate(ctx *gin.Context) {
	var exchangeRate model.ExchangeRate
	if err := ctx.ShouldBindJSON(&exchangeRate); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	exchangeRate.Date = time.Now()
	if err := global.Db.AutoMigrate(&exchangeRate); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := global.Db.Create(&exchangeRate).First(&exchangeRate).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"exchange_rate": exchangeRate,
	})
}

func GetExchangeRate(ctx *gin.Context) {
	var exchangeRates []model.ExchangeRate
	if err := global.Db.Find(&exchangeRates).Error; err != nil {
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
			"exchange_rates": exchangeRates,
		},
	)
}
