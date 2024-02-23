package main

import (
	"api.service.go/go-api-service/database"
)

func main() {
	database.OpenDB("./test.db")
}
