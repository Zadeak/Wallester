package unitTest

import (
	mapper "W/repository"
	repositorium "W/repository"
	"testing"
)

var repo = repositorium.DbCustomerRepo{}

func TestRepository(t *testing.T) {
	testCustomer, randomString := InitRandomCustomer(&repo)
	customer, _ := repo.InsertCustomer(&testCustomer)
	AssertEquals("Testing_Name"+randomString, customer.FirstName, t)
	DeleteCustomer(&customer, &repo)
}

func TestInsertCustomer_exists(t *testing.T) {
	testCustomer, _ := InitRandomCustomer(&repo)
	customer, _ := repo.InsertCustomer(&testCustomer)
	_, err := repo.InsertCustomer(&testCustomer)
	AssertEquals("ERROR: Customer is already registered", err.Error(), t)
	DeleteCustomer(&customer, &repo)
}

func TestUpdateCustomer(t *testing.T) {
	testCustomer, _ := InitRandomCustomer(&repo)
	customerInserted, _ := repo.InsertCustomer(&testCustomer)
	customerUpdated, _ := repo.UpdateCustomer(mapper.CustomerMapper(customerInserted))
	AssertEquals(customerUpdated.FirstName, customerInserted.FirstName, t)
	DeleteCustomer(&customerUpdated, &repo)
}

func TestUpdateCustomer_notFound(t *testing.T) {
	testCustomer, _ := InitRandomCustomer(&repo)
	_, err := repo.UpdateCustomer(&testCustomer)
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
	//
	customers, _ := repo.FindCustomers(&testCustomer1)
	//

	AssertListContainsCustomer(insertedCustomer1, customers, t)
	AssertListContainsCustomer(insertedCustomer2, customers, t)
	AssertListContainsCustomer(insertedCustomer3, customers, t)
	AssertListContainsCustomer(insertedCustomer4, customers, t)

	DeleteAllList(customers, &repo)
}
