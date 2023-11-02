package schemas

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Transaction struct {
	gorm.Model
	ID                 uuid.UUID
	TransactionValue   float64
	ProductDescription string
	CardNumber         float64
	NameInCard         string
	CardExpirationDate time.Time
	Cvv                int8
}
