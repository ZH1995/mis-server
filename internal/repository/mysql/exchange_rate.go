package mysql

import (
	"MyGin/global"
	"MyGin/internal/model"
)

func CreateExchangeRate(exchange_rate *model.ExchangeRate) error {
	return global.Db.Create(exchange_rate).Error
}

func GetExchangeRates() ([]model.ExchangeRate, error) {
	var exchange_rates []model.ExchangeRate
	err := global.Db.Find(&exchange_rates).Error
	return exchange_rates, err
}
