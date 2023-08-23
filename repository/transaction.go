package repository

import (
	"bpjs/model"
	"context"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	Create(ctx context.Context, db *gorm.DB, req []model.Transaction) error
}

type TransactionRepositoryImpl struct {
}

func NewTransactionRepository() TransactionRepository {
	return &TransactionRepositoryImpl{}
}

func (transactionRepository TransactionRepositoryImpl) Create(ctx context.Context, db *gorm.DB, req []model.Transaction) error {
	result := db.Create(&req)
	return result.Error
}
