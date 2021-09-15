package unitTest

import (
	mapper "W/repository"
	repositorium "W/repository"
	"testing"
)

// Repo tests
var repo = repositorium.DbCustomerRepo{}

func TestInsertCustomer(t *testing.T) {
	testCustomer, randomString := InitRandomCustomer(&repo)
	customer, _ := repo.InsertCustomer(&testCustomer)
	AssertEquals("Testing_Name"+randomString, customer.FirstName, t)
}

func TestInsertCustomer_exists(t *testing.T) {
	testCustomer, _ := InitRandomCustomer(&repo)
	_, _ = repo.InsertCustomer(&testCustomer)
	_, err := repo.InsertCustomer(&testCustomer)
	AssertEquals("ERROR: Customer is already registered", err.Error(), t)
}

func TestUpdateCustomer(t *testing.T) {
	testCustomer, _ := InitRandomCustomer(&repo)
	customerInserted, _ := repo.InsertCustomer(&testCustomer)
	customerUpdated, _ := repo.UpdateCustomer(mapper.CustomerMapper(customerInserted), int(customerInserted.ID))
	AssertEquals(customerUpdated.FirstName, customerInserted.FirstName, t)
}

func TestUpdateCustomer_notFound(t *testing.T) {
	testCustomer, _ := InitRandomCustomer(&repo)
	_, err := repo.UpdateCustomer(&testCustomer, 13214125)
	AssertEquals("ERROR: customer is not found", err.Error(), t)
}

func TestFindCustomers(t *testing.T) {
	testCustomer1, _ := InitConcreteCustomer(&repo)
	testCustomer2, _ := InitConcreteCustomer(&repo)
	testCustomer3, _ := InitConcreteCustomer(&repo)
	testCustomer4, _ := InitConcreteCustomer(&repo)

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
