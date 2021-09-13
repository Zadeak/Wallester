package types

import (
	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model

	FirstName string
	LastName  string
	Email     string `gorm:"typevarchar(100);unique_index"`
}

type CustomePogo struct { //Plain old Go Object
	FirstName string
	LastName  string
	Email     string
}
