package main

import (
	"lab3/Server"
)

func main() {
	var server = Server.NewServer()
	server.ListenAndServe()
}
