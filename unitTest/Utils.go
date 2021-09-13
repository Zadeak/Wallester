package unitTest

import (
	types "W/repository"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func AssertEquals(expected interface{}, actual interface{}, t *testing.T) {
	if expected != actual {
		t.Fatalf("expected is %v, actual is %v", expected, actual)
	}
}

func AssertListContainsCustomer(e types.Customer, s []types.Customer, t *testing.T) bool {
	for _, a := range s {
		if a.Email == e.Email {
			return true
		}
	}
	t.Fatalf("List does not contain %v", e)
	return false
}
func InitRandomCustomer(repo *types.DbCustomerRepo) (types.CustomerPogo, string) {
	rand.Seed(time.Now().UnixNano())
	s1 := strconv.FormatInt(rand.Int63(), 10)
	repo.Migrate()
	testCustomer := types.CustomerPogo{
		FirstName: "Testing_Name" + s1,
		LastName:  "Testing_LastName" + s1,
		Email:     "Email@unitTest" + s1,
	}
	return testCustomer, s1
}
func InitConcreteCustomer(repo *types.DbCustomerRepo) (types.CustomerPogo, string) {
	rand.Seed(time.Now().UnixNano())
	s1 := strconv.FormatInt(rand.Int63(), 10)
	repo.Migrate()
	testCustomer := types.CustomerPogo{
		FirstName: "Testing_Name",
		LastName:  "Testing_LastName",
		Email:     "Email@unitTest" + s1,
	}
	return testCustomer, s1
}
func DeleteCustomer(customer *types.Customer, repo *types.DbCustomerRepo) {
	err := repo.DeleteCustomer(&types.CustomerPogo{
		FirstName: customer.FirstName,
		LastName:  customer.LastName,
		Email:     customer.Email,
	})
	if err != nil {
		panic(err)
	}
}

func DeleteAll(customer []types.Customer, repo *types.DbCustomerRepo) {

	for i := range customer {
		t := customer[i]
		err := repo.DeleteCustomer(&types.CustomerPogo{
			FirstName: t.FirstName,
			LastName:  t.LastName,
			Email:     t.Email,
		})
		if err != nil {
			panic(err)
		}
	}

}
