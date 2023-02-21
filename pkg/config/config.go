package config

import (
	"fmt"
	"os"
)
/*
This implementation creates a Config type that has fields for the DatabaseURL and Port values. 
The NewConfig() function is used to load these values from environment variables, 
with DATABASE_URL being required and PORT being optional with a default value of 8080. 
If any of the required environment variables are missing, the function returns an error. 
Otherwise, it returns a pointer to a Config instance with the loaded configuration values. 
These values can then be used throughout the application.
*/
type Config struct {
	DatabaseURL string
	Port        string
}

func NewConfig() (*Config, error) {
	dbURL, ok := os.LookupEnv("DATABASE_URL")
	if !ok {
		return nil, fmt.Errorf("missing required environment variable DATABASE_URL")
	}

	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "8080"
	}

	return &Config{
		DatabaseURL: dbURL,
		Port:        port,
	}, nil
}
