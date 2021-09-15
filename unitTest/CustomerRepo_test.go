package unitTest

import (
	mapper "W/repository"
	repositorium "W/repository"
	"testing"
)

// Repo tests
var repo = repositorium.DbCustomerRepo{}

func TestInsertCustomer(t *testing.T) {
	initRepo(&repo)
	testCustomer, randomString := InitRandomCustomer()
	customer, _ := repo.InsertCustomer(&testCustomer)
	AssertEquals("Testing_Name"+randomString, customer.FirstName, t)
}

func TestInsertCustomer_exists(t *testing.T) {
	initRepo(&repo)
	testCustomer, _ := InitRandomCustomer()
	_, _ = repo.InsertCustomer(&testCustomer)
	_, err := repo.InsertCustomer(&testCustomer)
	AssertEquals("ERROR: Customer is already registered", err.Error(), t)
}

func TestUpdateCustomer(t *testing.T) {
	initRepo(&repo)
	testCustomer, _ := InitRandomCustomer()
	customerInserted, _ := repo.InsertCustomer(&testCustomer)
	customerUpdated, _ := repo.UpdateCustomer(mapper.CustomerMapper(customerInserted), int(customerInserted.ID))
	AssertEquals(customerUpdated.FirstName, customerInserted.FirstName, t)
}

func TestUpdateCustomer_notFound(t *testing.T) {
	initRepo(&repo)
	testCustomer, _ := InitRandomCustomer()
	_, err := repo.UpdateCustomer(&testCustomer, 13214125)
	AssertEquals("ERROR: customer is not found", err.Error(), t)
}

func TestFindCustomers(t *testing.T) {
	initRepo(&repo)
	testCustomer1, _ := InitConcreteCustomer()
	testCustomer2, _ := InitConcreteCustomer()
	testCustomer3, _ := InitConcreteCustomer()
	testCustomer4, _ := InitConcreteCustomer()

	insertedCustomer1, _ := repo.InsertCustomer(&testCustomer1)
	insertedCustomer2, _ := repo.InsertCustomer(&testCustomer2)
	insertedCustomer3, _ := repo.InsertCustomer(&testCustomer3)
	insertedCustomer4, _ := repo.InsertCustomer(&testCustomer4)

	customers, _ := repo.FindCustomers(&testCustomer1)

	AssertListContainsCustomer(insertedCustomer1, customers, t)
	AssertListContainsCustomer(insertedCustomer2, customers, t)
	AssertListContainsCustomer(insertedCustomer3, customers, t)
	AssertListContainsCustomer(insertedCustomer4, customers, t)
}
func TestFindCustomer(t *testing.T) {
	initRepo(&repo)
	testCustomer1, _ := InitConcreteCustomer()
	insertedCustomer1, _ := repo.InsertCustomer(&testCustomer1)

	customer, _ := repo.FindCustomer(int(insertedCustomer1.ID))

	AssertEquals(testCustomer1.FirstName, customer.FirstName, t)
	AssertEquals(testCustomer1.LastName, customer.LastName, t)
	AssertEquals(testCustomer1.Email, customer.Email, t)

}
