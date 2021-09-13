package main

import (
	L "W/repository"
	"W/types"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model

	FirstName string
	LastName  string
	Email     string `gorm:"typevarchar(100);unique_index"`
}

var db *gorm.DB
var err error

func main() {

	// db connection variables
	dbHost := "localhost"
	dbPort := 5432
	dbUser := "postgres"
	dbPass := "postgres"
	dbName := "postgres"
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPass, dbName)

	// Opening connection

	db, err = gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}
	fmt.Print(db)

	fmt.Println("Hello!")

	//make migrations to the database if the have not already beed created

	err := db.AutoMigrate(&Customer{})
	if err != nil {
		panic(err.Error())
	}

	customer := types.CustomePogo{
		FirstName: "Test",
		LastName:  "TestLastName",
		Email:     "emailTest@gmail.com",
	}

	insertedCustomer := L.Insert(&customer)

	println(insertedCustomer.ID)

	L.Update(&customer)

	//db.Create(Customer{FirstName: "Inserted from code name",LastName: "Inserted from code lastNmae", Email: "insertedfromcode@gmail.com"})

}
