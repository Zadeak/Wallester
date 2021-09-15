package controller

import (
	"W/repository"
	"W/server"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"strconv"
)

type Controller struct {
	Server *server.Server
	Repo   *repository.DbCustomerRepo
}

func (c *Controller) InitController() {
	c.Server = &server.Server{}
	c.Repo = &repository.DbCustomerRepo{}
}
func (c *Controller) InitRepo() {
	c.Repo.Initialize()
}

func (c *Controller) InitServer() {
	c.Server.InitRouter()
	c.InitiateRoutes()
	c.Server.StartServer()
}

func (c *Controller) InitiateRoutes() {
	c.Server.Router.HandleFunc("/api/search", c.searchCustomers)    //ok
	c.Server.Router.HandleFunc("/api/id={id}", c.showCustomer)      // ok
	c.Server.Router.HandleFunc("/api/create", c.createCustomer)     // ok
	c.Server.Router.HandleFunc("/api/id={id}/edit", c.editCustomer) //ok
}

func (c *Controller) createCustomer(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case "GET":
		http.ServeFile(w, r, "views/customerForm.gohtml")
	case "POST":
		c.processForm(w, r)
	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}

}

func (c *Controller) searchCustomers(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case "GET":
		http.ServeFile(w, r, "views/searchCustomersForm.html")
	case "POST":
		c.processSearchForm(w, r)
	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}
func (c *Controller) showCustomer(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Customer ID")
		return
	}

	t, err := template.New("showCustomer.gohtml").ParseFiles("views/showCustomer.gohtml")

	customer, err := c.Repo.FindCustomer(id)
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return

	}
	if err = t.Execute(w, &customer); err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}

}

func (c *Controller) editCustomer(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			respondWithError(w, http.StatusBadRequest, "Invalid Customer ID")
			return
		}
		t, err := template.New("editCustomer.gohtml").ParseFiles("views/editCustomer.gohtml")
		customer, _ := c.Repo.FindCustomer(id)
		if err = t.Execute(w, customer); err != nil {
			fmt.Println(err)
		}
	case "POST":
		vars := mux.Vars(r)
		id, _ := strconv.Atoi(vars["id"])
		customer, _ := c.Repo.FindCustomer(id)
		c.processFormUpdate(w, r, &customer)
	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}

}

func (c *Controller) processForm(w http.ResponseWriter, r *http.Request) {
	if checkForm(w, r) {
		return
	}
	name := r.FormValue("fname")
	lastName := r.FormValue("lname")
	email := r.FormValue("email")

	_, err := c.Repo.InsertCustomer(&repository.CustomerPogo{
		FirstName: name,
		LastName:  lastName,
		Email:     email,
	})

	if err != nil {
		// :TODO implement validation
		fmt.Fprintf(w, "ParseForm() err: %v", err)
	}
	http.ServeFile(w, r, "views/submitSuccess.gohtml")
}

func (c *Controller) processSearchForm(w http.ResponseWriter, r *http.Request) {
	if checkForm(w, r) {
		return
	}

	customers, err := c.Repo.FindCustomers(&repository.CustomerPogo{
		FirstName: r.FormValue("fname"),
		LastName:  r.FormValue("lname"),
		Email:     "",
	})

	if err != nil {
		//:TODO
	}
	t, err := template.New("showAllCustomers.gohtml").ParseFiles("views/showAllCustomers.gohtml")

	if err = t.Execute(w, map[string][]repository.Customer{"customers": customers}); err != nil {
		fmt.Println(err)
	}

}

func (c *Controller) processFormUpdate(w http.ResponseWriter, r *http.Request, customer *repository.Customer) {
	if checkForm(w, r) {
		return
	}

	pogo := repository.CustomerPogo{
		FirstName: r.FormValue("fname"),
		LastName:  r.FormValue("lname"),
		Email:     r.FormValue("email"),
	}
	_, err := c.Repo.UpdateCustomer(&pogo, int(customer.ID))
	if err != nil {
		fmt.Println(err)
	}
	http.ServeFile(w, r, "views/updateSuccess.gohtml")

}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
func checkForm(w http.ResponseWriter, r *http.Request) bool {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return true
	}
	return false
}

//t, err := template.New("customerForm.gohtml").ParseFiles("views/customerForm.gohtml")
//
//customers, _ := c.Repo.ShowCustomers()
//if err = t.Execute(w, map[string][]repository.Customer{"customers": customers}); err != nil {
//	fmt.Println(err)
//}
