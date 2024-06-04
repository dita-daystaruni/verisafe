package api

import (
	"fmt"

	"github.com/dita-daystaruni/verisafe/configs"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Server struct {
	*gin.Engine
	*configs.Config
	*gorm.DB
}

// Create a new server and load the configuration .env
func NewServer() (*Server, error) {
	// Load the configuration
	cfg, err := configs.LoadConfig()
	if err != nil {
		return nil, err
	}

	// Open the database connection
	con, err := connectToPostgres(cfg)
	if err != nil {
		return nil, err
	}

	// Instanciate gin's server engine
	server := gin.Default()

	return &Server{server, cfg, con}, nil
}

// binds and runs the server additionally it sets the various cors headers
func (s *Server) RunServer() {
	RegisterHandlers(s)
	s.Run(fmt.Sprintf(":%d", s.AppConfig.Port))
}

func connectToPostgres(cfg *configs.Config) (*gorm.DB, error) {
	dns := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s",
		cfg.DatabaseConfig.DatabaseUser,
		cfg.DatabaseConfig.DatabasePassword,
		cfg.DatabaseConfig.DatabaseHost,
		cfg.DatabaseConfig.DatabasePort,
		cfg.DatabaseConfig.DatabaseName,
	)

	// Attempt to open a connection
	con, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Return the db instance
	return con, nil
}
