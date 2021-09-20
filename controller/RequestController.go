package controller

import (
	"W/repository"
	"W/server"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"log"
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
}
func (c *Controller) StartServer() {
	c.Server.StartServer()
}

func (c *Controller) InitiateRoutes() {
	c.Server.Router.HandleFunc("/api/search", c.StartPaginationHandler).Methods("GET", "POST") //ok
	c.Server.Router.HandleFunc("/api/search/page={id}", c.ContinuePaginationHandler)           // ok
	c.Server.Router.HandleFunc("/api/create", c.CreateCustomer)                                // ok
	c.Server.Router.HandleFunc("/api/id={id}/edit", c.EditCustomer)                            //ok
	c.Server.Router.HandleFunc("/api/id={id}", c.ShowCustomer)
}

func (c *Controller) CreateCustomer(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case "GET":
		http.ServeFile(w, r, "views/customerForm.gohtml")
	case "POST":
		c.processForm(w, r)
	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}

}

func (c *Controller) ShowCustomer(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Customer ID")
	}

	t, err := template.New("showCustomer.gohtml").ParseFiles("views/showCustomer.gohtml")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	customer, err := c.Repo.FindCustomer(id)
	if err != nil {
		fmt.Fprintf(w, err.Error())

	}
	if err = t.Execute(w, customer); err != nil {
		fmt.Fprintf(w, err.Error())
	}

}

func (c *Controller) EditCustomer(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			respondWithError(w, http.StatusBadRequest, "Invalid Customer ID")
			return
		}
		t, err := template.New("EditCustomer.gohtml").ParseFiles("views/EditCustomer.gohtml")
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

func (c *Controller) ContinuePaginationHandler(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	page, _ := vars["id"]
	intVar, _ := strconv.Atoi(page)
	name := request.URL.Query().Get("name")
	lastName := request.URL.Query().Get("last-name")

	list, _ := c.Repo.List(repository.Pagination{
		Page:     intVar,
		Name:     name,
		LastName: lastName,
	}, &repository.CustomerPogo{
		FirstName: name,
		LastName:  lastName,
	})
	tpl := template.Must(template.ParseGlob("views/*.gohtml"))
	if err := tpl.ExecuteTemplate(writer, "pagination.gohtml", list); err != nil {
		log.Println(err)
	}
	return
}

func (c *Controller) StartPaginationHandler(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {

	case "GET":
		http.ServeFile(writer, request, "views/searchCustomersForm.html")
	case "POST":
		displayFirstSearchPage(writer, request, c)
		return

	default:
		fmt.Fprintf(writer, "Sorry, only GET and POST methods are supported.")
	}
}

func displayFirstSearchPage(writer http.ResponseWriter, request *http.Request, c *Controller) {
	request.FormValue("fname")
	request.FormValue("lname")
	vars := mux.Vars(request)
	page, _ := vars["id"]
	intVar, _ := strconv.Atoi(page)

	nameValue := request.FormValue("fname")
	lastNameValue := request.FormValue("lname")

	list, _ := c.Repo.List(repository.Pagination{
		Page:     intVar,
		Name:     nameValue,
		LastName: lastNameValue,
	}, &repository.CustomerPogo{
		FirstName: nameValue,
		LastName:  lastNameValue,
	})
	tpl := template.Must(template.ParseGlob("views/*.gohtml"))
	if err := tpl.ExecuteTemplate(writer, "pagination.gohtml", list); err != nil {
		log.Println(err)
	}
	return
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
