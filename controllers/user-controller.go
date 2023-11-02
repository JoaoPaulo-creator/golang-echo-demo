package controllers

import (
	"example.io/echo-demo/config"
	"fmt"
	"net/http"

	"example.io/echo-demo/dtos"
	"example.io/echo-demo/utils"
	"github.com/labstack/echo/v4"
)

type ErrorResponse struct {
	message string `json:"message"`
}

func CreateUserController(c echo.Context) error {
	var transactionDTO dtos.TransactionDTO

	if err := c.Bind(&transactionDTO); err != nil {
		return err
	}

	if transactionDTO.TransactionValue != 0 {
		fmt.Println("VALOR DA TRANSACTION: ", transactionDTO.TransactionValue)
	}

	fmt.Print("NAME IN CARD: ", transactionDTO.NameInCard)

	transaction := utils.TransactionDTOConverter(transactionDTO)

	// TODO: identificar erro e corrigir

	if err := &config.GormDB.Create(&transaction).Error; err != nil {
		return c.JSON(http.StatusBadRequest, "Ocorreu um erro ao tentar criar transacao")
	}

	response := utils.TransactionResponseConverter(transaction)
	return c.JSON(http.StatusCreated, response)
}
