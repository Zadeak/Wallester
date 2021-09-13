package repository

import (
	"W/types"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DbConn() (db *gorm.DB) {
	dbHost := "localhost"
	dbPort := 5432
	dbUser := "postgres"
	dbPass := "postgres"
	dbName := "postgres"
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPass, dbName)

	db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}
	return db
}

//Insert customer into db

func Insert(customer *types.CustomePogo) types.Customer {
	customerToInsert := types.Customer{
		Model:     gorm.Model{},
		FirstName: customer.FirstName,
		LastName:  customer.LastName,
		Email:     customer.Email,
	}
	db := DbConn()
	db.Create(&customerToInsert)

	return customerToInsert
}

func Update(customer *types.CustomePogo) /*(types.Customer) */ {

	db := DbConn()

	var customers = types.Customer{}

	db.Find(&customers)

	fmt.Println(customers) //first := db.Where("customers.first_name = ?", customer.FirstName).First()

	//println(first)

	//if err != nil {
	//
	//	db.Save(&customer)
	//}

}
