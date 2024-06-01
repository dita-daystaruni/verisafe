// The db package provides primitives for database configuration and
// connection
package db

import (
	"fmt"
	"os"

	"github.com/dita-daystaruni/verisafe/internal/models/roles"
	"github.com/dita-daystaruni/verisafe/internal/models/user"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB = nil

func Connect() (*gorm.DB, error) {
	dns := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	// Attempt to open a connection
	con, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	DB = con

	// Return the db instance
	return DB, nil
}

func AutoMigrate() {
	DB.AutoMigrate(&roles.Role{}, &user.User{})
}
