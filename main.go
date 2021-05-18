package main

import server2 "github.com/abnerdsilva/webapi_mvc_with_go/server"

func main() {
	server := server2.NewServer()
	server.Run()
}
