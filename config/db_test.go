package config

import (
	"log"
	"os"
	"os/exec"
	"testing"

	"github.com/dita-daystaruni/verisafe/config/db"
)

func init() {
	exec.Command("createdb verisafe_test")
	os.Setenv("DB_USER", "erick")
	os.Setenv("DB_PASSWORD", "muuo")
	os.Setenv("DB_NAME", "verisafe_test")
	os.Setenv("DB_PORT", "5432")
	os.Setenv("DB_HOST", "localhost")

	log.Println("Test setup complete")
}

// Test database connection functionality
func TestDatabaseConnection(t *testing.T) {
	con, err := db.Connect()
	if err != nil {
		t.Fatalf("Failed to connect to database because: %s\n", err.Error())
	}

	if con == nil {
		t.Fatal("Database instance is null!")
	}
}
