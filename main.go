package main

import (
	"github.com/abnerdsilva/webapi_mvc_with_go/database"
	"github.com/abnerdsilva/webapi_mvc_with_go/server"
)

func main() {
	database.StartDB()

	server := server.NewServer()
	server.Run()
}
