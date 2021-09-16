package App

import (
	"W/controller"
)

type App struct {
	Controller *controller.Controller
}

func (a *App) Initialize() {
	a.start()

}

func (a *App) start() {
	a.Controller = &controller.Controller{}
	a.Controller.InitController()
	a.Controller.InitRepo()
	a.Controller.InitServer()
	a.Controller.StartServer()
}
