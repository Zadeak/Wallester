package repository

import (
	"W/types"
	_ "W/types"
)

type RepoInterface interface {
	InsertCustomer() types.Customer
	UpdateCustomer()
	QueryCustomer()
}
