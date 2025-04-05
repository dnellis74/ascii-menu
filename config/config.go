package config

import (
	"embed"
	"encoding/json"
	"os"
)

//go:embed games/config.json
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

// LoadConfig loads both embedded and external game configurations
func LoadConfig() (*Config, error) {
	// Load embedded config
	embeddedConfig, err := loadEmbeddedConfig()
	if err != nil {
		return nil, err
	}

	// Try to load external config
	externalConfig, err := loadExternalConfig()
	if err != nil {
		// If external config doesn't exist, just return embedded config
		if os.IsNotExist(err) {
			return embeddedConfig, nil
		}
		return nil, err
	}

	// Merge configurations
	mergedConfig := &Config{
		Games: append(embeddedConfig.Games, externalConfig.Games...),
	}

	return mergedConfig, nil
}

func loadEmbeddedConfig() (*Config, error) {
	file, err := configFS.ReadFile("games/config.json")
	if err != nil {
		return nil, err
	}

	var config Config
	if err := json.Unmarshal(file, &config); err != nil {
		return nil, err
	}

	return &config, nil
}

func loadExternalConfig() (*Config, error) {
	// Standard Linux config location
	configPath := "/etc/ascii-menu/config.json"

	file, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	var config Config
	if err := json.Unmarshal(file, &config); err != nil {
		return nil, err
	}

	return &config, nil
}
