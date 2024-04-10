package main

import (
	"api.service.go/go-api-service/app"
)

func main() {
	app := app.NewApp("10001", "./test.db")
	app.InitApp()
}
