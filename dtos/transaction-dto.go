package dtos

import (
	"github.com/google/uuid"
	"time"
)

type TransactionDTO struct {
	TransactionValue   float64   `json:"transaction_value"`
	ProductDescription string    `json:"product_description"`
	CardNumber         float64   `json:"card_number"`
	NameInCard         string    `json:"name_in_card"`
	CardExpirationDate time.Time `json:"card_expiration_date"`
	Cvv                int8      `json:"cvv"`
}

type TransactionResponse struct {
	ID                 uuid.UUID `json:"id"`
	TransactionValue   float64   `json:"transaction_value"`
	ProductDescription string    `json:"product_description"`
	CardNumber         float64   `json:"card_number"`
	NameInCard         string    `json:"name_in_card"`
	CardExpirationDate time.Time `json:"card_expiration_date"`
	Cvv                int8      `json:"cvv"`
}
