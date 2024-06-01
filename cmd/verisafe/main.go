package main

import (
	_ "github.com/dita-daystaruni/verisafe/config" // Side effect for loafing env vars
	"github.com/dita-daystaruni/verisafe/internal/handlers"
)

func main() {
	handlers.Serve()
}
