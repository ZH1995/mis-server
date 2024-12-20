package v1

import (
	"MyGin/internal/model"
	"MyGin/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateExchangeRate(ctx *gin.Context) {
	var exchangeRate model.ExchangeRate
	if err := ctx.ShouldBindJSON(&exchangeRate); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := service.CreateExchangeRate(&exchangeRate); err != nil {
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
	exchangeRates, err := service.GetExchangeRates()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(
		http.StatusOK,
		gin.H{
			"exchange_rates": exchangeRates,
		},
	)
}
