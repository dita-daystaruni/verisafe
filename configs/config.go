package configs

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/charmbracelet/log"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	// JWT token configuration
	JWTConfig struct {
		ApiSecret   string `envconfig:"API_SECRET"`
		ExpireDelta int    `envconfig:"EXPIRE_DELTA"`
	}

	// Application configuration
	AppConfig struct {
		Port   int    `envconfig:"VERISAFE_PORT"`
		Addres string `envconfig:"VERISAFE_ADDRESS"`
	}

	// Database configuration
	DatabaseConfig struct {
		DatabaseHost     string `envconfig:"DB_HOST"`
		DatabaseDriver   string `envconfig:"DB_DRIVER"`
		DatabaseUser     string `envconfig:"DB_USER"`
		DatabasePassword string `envconfig:"DB_PASSWORD"`
		DatabaseName     string `envconfig:"DB_NAME"`
		DatabasePort     int    `envconfig:"DB_PORT"`
	}

	APISecrets struct {
		EventApiSecret string `envconfig:"API_KEY"`
	}

	EventConfig struct {
		UserCreateEvent  []string `json:"user-create-event"`
		UserUpdatedEvent []string `json:"user-update-event"`
		UserDeletedEvent []string `json:"user-delete-event"`
	}

	Logger *log.Logger
}

// The LoadConfig function loads the env file specified and returns
// a valid configuration object ready for use
func LoadConfig() (*Config, error) {
	cfg := Config{}

	// Init the file logger
	file, err := os.OpenFile("error.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Default().Fatalf("Failed to open log file: %v", err)
	}
	cfg.Logger = log.NewWithOptions(file, log.Options{ReportTimestamp: true,
		ReportCaller: true,
	})

	// load the configs
	if err := godotenv.Load(".env"); err != nil {
		return nil, fmt.Errorf("Failed to load environment variables: %v", err)
	}
	if err := envconfig.Process("", &cfg); err != nil {
		return nil, fmt.Errorf("Failed to load environment variables: %v", err)
	}

	// Load the service configuration
	// Load JSON configuration file
	jsonFile, err := os.ReadFile("services.json")
	if err != nil {
		cfg.Logger.Error("Failed to read services.json", err)
		return nil, fmt.Errorf("Failed to read JSON config file: %v", err)
	}

	if err := json.Unmarshal(jsonFile, &cfg.EventConfig); err != nil {
		cfg.Logger.Error("Failed to read services.json", err)
		return nil, fmt.Errorf("Failed to unmarshal JSON config file: %v", err)
	}

	cfg.Logger.Info("Configuration files loaded successfully!")
	return &cfg, nil
}
