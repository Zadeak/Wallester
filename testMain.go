package main

import types "W/repository"

func main() {

	repo := types.NewCustomerRepo()
	repo.Migrate()
	//
	//
	customerToDelete := types.CustomerPogo{
		FirstName: "",
		LastName:  "",
		Email:     "test66",
	}
	repo.DeleteCustomer(&customerToDelete)

	//var newPogo = types.CustomerPogo{
	//	FirstName: "TestingUpdated2",
	//	LastName:  "Test",
	//	Email:     "test66",
	//}

	//insertedUser, err := repo.InsertCustomer(&newPogo)
	//if err != nil {
	//	panic(err)
	//}
	//println(insertedUser.ID)

	/*	customers, _ := repo.FindCustomers(&newPogo)

		println("customers: ",customers)

		for _, customer := range customers {
			customer.ToString()
		}

		for i := range customers {
			println(customers[i].ToString())
		}*/

	//update
	/*	customer, _ := repo.UpdateCustomer(&newPogo)

		println(customer.Email)*/

	//insertedUser, err := repo.InsertCustomer(&newPogo)
	//if err != nil {
	//	panic(err)
	//}
	//println(insertedUser.ID)

}
