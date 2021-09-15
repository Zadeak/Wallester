package unitTest

import (
	repositorium "W/repository"
	"math/rand"
	"os"
	"strconv"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	var repo = repositorium.DbCustomerRepo{}
	ensureTableExists(&repo)
	code := m.Run()
	clearTable(&repo)
	os.Exit(code)
}

func ensureTableExists(repo *repositorium.DbCustomerRepo) {
	repo.Initialize()

}
func clearTable(repo *repositorium.DbCustomerRepo) {
	db, _ := repo.Db.DB()
	db.Exec("DELETE FROM customers")
	db.Exec("ALTER SEQUENCE customers_id_seq RESTART WITH 1")
}

func AssertEquals(expected interface{}, actual interface{}, t *testing.T) {
	if expected != actual {
		t.Fatalf("expected is %v, actual is %v", expected, actual)
	}
}

func AssertListContainsCustomer(e repositorium.Customer, s []repositorium.Customer, t *testing.T) bool {
	for _, a := range s {
		if a.Email == e.Email {
			return true
		}
	}
	t.Fatalf("List does not contain %v", e)
	return false
}
func InitRandomCustomer(repo *repositorium.DbCustomerRepo) (repositorium.CustomerPogo, string) {
	rand.Seed(time.Now().UnixNano())
	s1 := strconv.FormatInt(rand.Int63(), 10)
	repo.Initialize()
	testCustomer := repositorium.CustomerPogo{
		FirstName: "Testing_Name" + s1,
		LastName:  "Testing_LastName" + s1,
		Email:     "Email@unitTest" + s1,
	}
	return testCustomer, s1
}
func InitConcreteCustomer(repo *repositorium.DbCustomerRepo) (repositorium.CustomerPogo, string) {
	rand.Seed(time.Now().UnixNano())
	s1 := strconv.FormatInt(rand.Int63(), 10)
	repo.Initialize()
	testCustomer := repositorium.CustomerPogo{
		FirstName: "Testing_Name",
		LastName:  "Testing_LastName",
		Email:     "Email@unitTest" + s1,
	}
	return testCustomer, s1
}
func DeleteCustomer(customer *repositorium.Customer, repo *repositorium.DbCustomerRepo) {
	err := repo.DeleteCustomer(&repositorium.CustomerPogo{
		FirstName: customer.FirstName,
		LastName:  customer.LastName,
		Email:     customer.Email,
	})
	if err != nil {
		panic(err)
	}
}

func DeleteAllList(customer []repositorium.Customer, repo *repositorium.DbCustomerRepo) {

	for i := range customer {
		t := customer[i]
		err := repo.DeleteCustomer(&repositorium.CustomerPogo{
			FirstName: t.FirstName,
			LastName:  t.LastName,
			Email:     t.Email,
		})
		if err != nil {
			panic(err)
		}
	}

}

/*func TestMain(m *testing.M, repo *repositorium.DbCustomerRepo) {
	ensureTableExists(repo)
	code := m.Run()
	clearTable(repo)
	os.Exit(code)
}

func ensureTableExists(repo *repositorium.DbCustomerRepo) {
	repo.Initialize()

}
func clearTable(repo *repositorium.DbCustomerRepo) {
	db, _ := repo.Db.DB()
	db.Exec("DELETE FROM postgres")
	db.Exec("ALTER SEQUENCE customers_id_seq RESTART WITH 1")
}*/
