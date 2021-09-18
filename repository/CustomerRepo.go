package repository

import (
	"errors"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type CustomerPogo struct { //Plain old Go Object
	FirstName string
	LastName  string
	Email     string
}
type Customer struct {
	gorm.Model

	FirstName string
	LastName  string
	Email     string `gorm:"typevarchar(100);unique_index;unique"`
}

type DbCustomerRepo struct {
	Db *gorm.DB
}

func (repo *DbCustomerRepo) Initialize() {
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
	repo.Db = db
	repo.migrate()
}
func (repo *DbCustomerRepo) migrate() {
	err := repo.Db.AutoMigrate(&Customer{})
	if err != nil {
		panic(err.Error())
	}

}

func (repo *DbCustomerRepo) Migrate() {
	err := repo.Db.AutoMigrate(&Customer{})
	if err != nil {
		return
	}

}
func (repo *DbCustomerRepo) InsertCustomer(customerPogo *CustomerPogo) (Customer, error) {

	_, err := repo.findCustomerByEmail(customerPogo)

	if err != nil {
		newCustomer := Customer{
			Model:     gorm.Model{},
			FirstName: customerPogo.FirstName,
			LastName:  customerPogo.LastName,
			Email:     customerPogo.Email,
		}
		repo.Db.Create(&newCustomer)
		return newCustomer, nil
	} else {
		return Customer{}, errors.New("ERROR: Customer is already registered")
	}
}
func (repo *DbCustomerRepo) FindCustomers(customerPogo *CustomerPogo) ([]Customer, error) {
	var customers []Customer
	repo.Db.Where(Customer{FirstName: customerPogo.FirstName}).Or(Customer{LastName: customerPogo.LastName}).Find(&customers)
	return customers, nil
}
func (repo *DbCustomerRepo) FindCustomer(customerData int) (Customer, error) {
	var customer Customer
	repo.Db.Find(&customer, customerData)
	if customer.ID == 0 {
		return Customer{}, errors.New("ERROR: customer is not found")
	}
	return customer, nil
}

func (repo *DbCustomerRepo) UpdateCustomer(customerPogo *CustomerPogo, customerId int) (Customer, error) {
	customer := Customer{}
	repo.findCustomerById(customerId, &customer)
	if customer.ID == 0 {
		return Customer{}, errors.New("ERROR: customer is not found")
	} else {
		repo.updateCustomer(customerPogo, &customer)
		return customer, nil
	}
}

func (repo *DbCustomerRepo) updateCustomer(customerPogo *CustomerPogo, customer *Customer) *Customer {

	repo.Db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&customer).Updates(map[string]interface{}{"first_name": customerPogo.FirstName, "last_name": customerPogo.LastName, "email": customerPogo.Email}).Error; err != nil {
			return err
		}
		return nil
	})
	return customer
}

func (repo *DbCustomerRepo) findCustomerByEmail(customerPogo *CustomerPogo) (Customer, error) {
	var customer Customer
	repo.Db.Where(CustomerPogo{Email: customerPogo.Email}).First(&customer)
	if customer.ID == 0 {
		return Customer{}, errors.New("ERROR: customer is not found")
	}
	return customer, nil

}

func (repo *DbCustomerRepo) findCustomerById(customerId int, customer *Customer) *gorm.DB {
	return repo.Db.First(&customer, customerId)

}
