package config

import (
	"encoding/json"
	"os"
)

// Game represents a single game entry in the configuration
type Game struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Path        string `json:"path"`
}

// Config represents the entire configuration file
type Config struct {
	Games []Game `json:"games"`
}

// LoadConfig loads the game configuration from a JSON file
func LoadConfig(path string) (*Config, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var config Config
	if err := json.Unmarshal(file, &config); err != nil {
		return nil, err
	}

	return &config, nil
}
