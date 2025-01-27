package main

import (
	"log"

	"github.com/dita-daystaruni/verisafe/api"
)

func main() {
	server, err := api.NewServer()
	if err != nil {
		log.Fatal(err.Error())
	}

	server.RunServer()
}
