package model

import (
	// d "github.com/shopspring/decimal"
)

type Transaction struct {
	Id         int       `json:"id"`
	CustomerId int       `json:"customer_id"`
	TranType   string    `json:"tran_type"`
	TranCode   string    `json:"tran_code"`
	TranDate   string    `json:"tran_date"`
	Amount     float64 `json:"amount"`
}
