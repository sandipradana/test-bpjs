package service

import (
	"bpjs/model"
	"bpjs/repository"
	"context"
	"fmt"

	"gorm.io/gorm"
)

type TransactionService interface {
	Create(ctx context.Context, req []model.Transaction) error
}

type TransactionServiceImpl struct {
	db                    *gorm.DB
	transactionRepository repository.TransactionRepository
}

func NewTransactionService(db *gorm.DB, transactionRepository repository.TransactionRepository) TransactionService {
	return &TransactionServiceImpl{
		db:                    db,
		transactionRepository: transactionRepository,
	}
}

func (transactionService TransactionServiceImpl) Create(ctx context.Context, req []model.Transaction) error {
	err := transactionService.transactionRepository.Create(ctx, transactionService.db, req)
	if err != nil {
		fmt.Println(err)
	}

	return err
}
