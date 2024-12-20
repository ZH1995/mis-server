package service

import (
	"MyGin/internal/model"
	"MyGin/internal/repository/mysql"
)

func GetExchangeRates() ([]model.ExchangeRate, error) {
	return mysql.GetExchangeRates()
}

func CreateExchangeRate(exchange_rate *model.ExchangeRate) error {
	return mysql.CreateExchangeRate(exchange_rate)
}
