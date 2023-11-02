package schemas

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID     uint
	Name   string
	Age    uint
	Salary float64
}
