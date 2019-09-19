package main

import (
	"rentcar/webserver"
)

func main() {
	//create a new GIN endpoint server
	server := webserver.New()
	server.Run()
}
