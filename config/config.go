package config

import (
	"embed"
	"encoding/json"
)

//go:embed config/games/config.json
var configFS embed.FS

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

// LoadConfig loads the game configuration from the embedded JSON file
func LoadConfig() (*Config, error) {
	file, err := configFS.ReadFile("config/games/config.json")
	if err != nil {
		return nil, err
	}

	var config Config
	if err := json.Unmarshal(file, &config); err != nil {
		return nil, err
	}

	return &config, nil
}
