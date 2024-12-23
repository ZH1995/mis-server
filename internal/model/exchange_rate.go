package model

type ExchangeRate struct {
	ID           uint    `json:"id" gorm:"primarykey"`
	FromCurrency string  `json:"from_currency" binding:"required"`
	ToCurrency   string  `json:"to_currency" binding:"required"`
	Rate         float64 `json:"rate" binding:"required"`
	Date         int64   `json:"date"`
}
