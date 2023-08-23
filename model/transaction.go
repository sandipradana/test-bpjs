package model

import "time"

type Transaction struct {
	Id        uint      `json:"id"`
	Customer  string    `json:"customer"`
	Price     float32   `json:"price"`
	Timestamp time.Time `json:"timestamp"`
}

type TransactionCreateRequest struct {
	RequestId uint                           `json:"request_id" binding:"required"`
	Data      []TransactionCreateDataRequest `json:"data" binding:"required,dive"`
}

type TransactionCreateDataRequest struct {
	Id        uint    `json:"id" binding:"required"`
	Customer  string  `json:"customer" binding:"required"`
	Price     float32 `json:"price" binding:"required"`
	Timestamp string  `json:"timestamp" binding:"required"`
}
