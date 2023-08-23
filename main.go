package main

import (
	"bpjs/controller"
	"bpjs/repository"
	"bpjs/service"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	dsn := "root@tcp(127.0.0.1:3306)/transaction?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	transactionRepository := repository.NewTransactionRepository()
	transactionService := service.NewTransactionService(db, transactionRepository)
	transactionController := controller.NewTransactionController(transactionService)

	r := gin.Default()

	r.POST("/transaction", transactionController.Create)

	err = r.Run()
	if err != nil {
		panic(err)
	}
}
