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
		return
	}

}
func (customer *Customer) ToString() string {
	return fmt.Sprintf("ID: %v \n First name: %v \n LastName: %v", customer.ID, customer.FirstName, customer.LastName)
}
func (repo *DbCustomerRepo) Migrate() {
	err := repo.Db.AutoMigrate(&Customer{})
	if err != nil {
		return
	}

}
func (repo *DbCustomerRepo) InsertCustomer(customerPogo *CustomerPogo) (Customer, error) {

	var customer Customer
	repo.findCustomer(customerPogo, &customer)

	if customer.ID == 0 {
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
	repo.Db.Where(map[string]interface{}{"first_name": customerPogo.FirstName, "last_name": customerPogo.LastName}).Find(&customers)
	return customers, nil
}

func (repo *DbCustomerRepo) ShowCustomers() ([]Customer, error) {
	var customers []Customer
	repo.Db.Find(&customers)
	return customers, nil
}

func (repo *DbCustomerRepo) DeleteCustomer(customerPogo *CustomerPogo) error {
	customer := Customer{}
	repo.findCustomer(customerPogo, &customer)
	if customer.ID == 0 {
		return errors.New("ERROR: customer is not found")
	} else {
		repo.Db.Delete(&customer)
		return nil
	}

}

func (repo *DbCustomerRepo) F() []Customer {
	var customers []Customer
	repo.Db.Find(&customers)

	return customers
}

func (repo *DbCustomerRepo) UpdateCustomer(customerPogo *CustomerPogo) (Customer, error) {
	customer := Customer{}
	repo.findCustomer(customerPogo, &customer)
	if customer.ID == 0 {
		return Customer{}, errors.New("ERROR: customer is not found")
	} else {
		repo.updateCustomer(customerPogo, customer)
		return customer, nil
	}
}

func (repo *DbCustomerRepo) updateCustomer(customerPogo *CustomerPogo, customer Customer) *gorm.DB {
	return repo.Db.Model(&customer).Updates(map[string]interface{}{"first_name": customerPogo.FirstName, "last_name": customerPogo.LastName})
}

func (repo *DbCustomerRepo) findCustomer(customerPogo *CustomerPogo, customer *Customer) *gorm.DB {
	return repo.Db.Where(CustomerPogo{Email: customerPogo.Email}).First(&customer)

}

func DbConnection() (db *gorm.DB) {
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
