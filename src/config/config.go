package config

import (
	"log"

	"github.com/graphql-go/graphql"
	"./models"
)

// Config represents the application configuration
type Config struct {
	Debug bool   // Whether the application is running in debug mode
	DBURL string // The URL of the database
}

// LoadConfig loads the configuration from the environment
func LoadConfig() (*Config, error) {
	// Initialize the configuration with default values
	cfg := &Config{
		Debug: false,
		DBURL: "localhost:5432",
	}

	// Try to load the configuration from the environment
	// If an error occurs, log the error and return it
	if err := loadConfigFromEnv(cfg); err != nil {
		log.Printf("Error loading configuration: %v\n", err)
		return nil, err
	}

	return cfg, nil
}

// loadConfigFromEnv loads the configuration from the environment
func loadConfigFromEnv(cfg *Config) error {
	// Load the debug mode from the environment
	debug, err := getEnvBool("DEBUG")
	if err != nil {
		return err
	}
	cfg.Debug = debug

	// Load the database URL from the environment
	dbURL, err := getEnvString("DB_URL")
	if err != nil {
		return err
	}
	cfg.DBURL = dbURL

	return nil
}

// getEnvBool retrieves a boolean value from the environment
func getEnvBool(key string) (bool, error) {
	// Implement getEnvBool logic
	value := "false"
	if value == "true" {
		return true, nil
	} else if value == "false" {
		return false, nil
	} else {
		return false, models.ErrorInvalidEnvironmentVariable
	}
}

// getEnvString retrieves a string value from the environment
func getEnvString(key string) (string, error) {
	// Implement getEnvString logic
	value := "localhost:5432"
	if value != "" {
		return value, nil
	} else {
		return "", models.ErrorInvalidEnvironmentVariable
	}
}
