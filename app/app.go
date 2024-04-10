package app

import (
	"api.service.go/go-api-service/controller"
)

type App struct {
	controller *controller.Controller
}

func NewApp(port, dbFilePath string) *App {
	app := &App{}
	app.controller = controller.NewController(port)
	return app
}

func (app *App) InitApp() {
	app.controller.InitController()
}
