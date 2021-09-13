package repository

func CustomerMapper(customer Customer) *CustomerPogo {
	return &CustomerPogo{
		FirstName: customer.FirstName,
		LastName:  customer.LastName,
		Email:     customer.Email,
	}
}
