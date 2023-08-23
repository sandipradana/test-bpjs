package controller

import (
	"bpjs/model"
	"bpjs/service"
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type TransactionController interface {
	Create(c *gin.Context)
}

type TransactionControllerImpl struct {
	transactionService service.TransactionService
}

func NewTransactionController(transactionService service.TransactionService) TransactionController {
	return &TransactionControllerImpl{
		transactionService: transactionService,
	}
}

func (transactionController TransactionControllerImpl) Create(c *gin.Context) {

	var transactionRequest model.TransactionCreateRequest
	var ctx context.Context

	if err := c.ShouldBindJSON(&transactionRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": err.Error()})
		return
	}

	var transactions []model.Transaction

	for i := range transactionRequest.Data {

		layout := "2006-01-02 15:04:05"
		parsedTime, _ := time.Parse(layout, transactionRequest.Data[i].Timestamp)

		transactions = append(transactions, model.Transaction{
			Id:        transactionRequest.Data[i].Id,
			Customer:  transactionRequest.Data[i].Customer,
			Price:     transactionRequest.Data[i].Price,
			Timestamp: parsedTime,
		})
	}

	err := transactionController.transactionService.Create(ctx, transactions)
	if err != nil {
		fmt.Println("request_id : ", transactionRequest.RequestId, "error message :", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": "success", "message": "Data created successfully"})

}
