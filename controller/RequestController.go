package controller

import (
	"W/repository"
	"W/server"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
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
	c.initiateRoutes()
	c.Server.StartServer()
}

func (c *Controller) initiateRoutes() {
	c.Server.Router.HandleFunc("/api", c.getCustomers)
}

func (c *Controller) getCustomers(w http.ResponseWriter, r *http.Request) {

	t, err := template.New("show.gohtml").ParseFiles("views/show.gohtml")
	customers, _ := c.Repo.ShowCustomers()

	err = t.Execute(w, map[string][]repository.Customer{"customers": customers})

	if err != nil {
		fmt.Println(err)
	}
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

//var customer repository.CustomerPogo
//
//if err := json.NewDecoder(r.Body).Decode(&customer); err != nil{
//
//}
//fmt.Fprintf(w,"<h1>Hello there</h1> <p>%v</p>",showCustomers)

//func(s *Server) initializeRoutes(){
//
//	s.Router.HandleFunc("/api/customers",s.getCustomers).Methods("Get")
//
//}
