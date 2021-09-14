package main

import (
	"W/controller"
)

type App struct {
	controller *controller.Controller
}

func (a *App) Initialize() {
	a.start()

}

func (a *App) start() {
	a.controller = &controller.Controller{}
	a.controller.InitController()
	a.controller.InitRepo()
	a.controller.InitServer()
}

//func(a *App) initializeRoutes(){
//	a.Server.AddHandler("/api/customer", a.getCustomers)
//	a.Server.AddHandler("/api/customer", a.setCustomers)
//	a.Server.AddHandler("/api/customer", a.removeCustomers)
//}
//
//func (a *App) getCustomers(w http.ResponseWriter, r *http.Request){
//
//
//}
//func (a *App) setCustomers(w http.ResponseWriter, r *http.Request){
//
//
//}
//func (a *App) removeCustomers(w http.ResponseWriter, r *http.Request){
//
//
//}
