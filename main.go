package main

import (
	"log"

	"github.com/dita-daystaruni/verisafe/internal/app"
)

func main() {
	server, err := app.NewServer()

	if err != nil {
		log.Fatal(err.Error())
	}

	server.RunServer()
}
