package main

import "RENTCAR/webserver"

func main() {
	// create a new Gin endpoint server
	server := webserver.New()
	server.Run()

}
