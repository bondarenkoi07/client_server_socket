package main

import "lab3/Client"

func main() {
	var client = Client.NewClient()
	client.ListenAndServe()
}
