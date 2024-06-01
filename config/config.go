// The config package loads the environment variables
// critical for the application to function
package config

import (
	"log"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to load .env file")
	}
	log.Printf("env variables loaded!\n")
}
