package utils

import (
	"example.io/echo-demo/dtos"
	"example.io/echo-demo/schemas"
	"github.com/google/uuid"
	"time"
)

func TransactionDTOConverter(t dtos.TransactionDTO) *schemas.Transaction {
	return &schemas.Transaction{
		ID:                 uuid.New(),
		TransactionValue:   t.TransactionValue,
		ProductDescription: t.ProductDescription,
		CardNumber:         t.CardNumber,
		NameInCard:         t.NameInCard,
		CardExpirationDate: time.Now(),
		Cvv:                t.Cvv,
	}
}

func TransactionResponseConverter(r *schemas.Transaction) dtos.TransactionResponse {
	return dtos.TransactionResponse{
		ID:                 r.ID,
		TransactionValue:   r.TransactionValue,
		ProductDescription: r.ProductDescription,
		CardNumber:         r.CardNumber,
		NameInCard:         r.NameInCard,
		CardExpirationDate: r.CardExpirationDate,
		Cvv:                r.Cvv,
	}
}
