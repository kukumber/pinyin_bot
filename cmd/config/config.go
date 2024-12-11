package config

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

// Config represents the configuration
type Config struct {
	APIKey         string
	InitOffset     int
	UpdateInterval int
}

// ParseFlags parses command line flags and returns a Config struct
func ParseFlags() (Config, error) {
	cfg := Config{}

	flag.StringVar(&cfg.APIKey, "api-key", "", "Telegram API key")
	flag.IntVar(&cfg.InitOffset, "init-offset", 0, "Initial Telegram offset")
	flag.IntVar(&cfg.UpdateInterval, "update-interval", 60, "Update interval in seconds")
	flag.Parse()

	if err := parseEnv(&cfg); err != nil {
		return cfg, err
	}

	if err := checkConfigValues(&cfg); err != nil {
		return cfg, err
	}

	return cfg, nil
}

// parseEnv parses environment variables and overrides the config values if empty
func parseEnv(cfg *Config) error {
	if cfg.APIKey == "" && os.Getenv("API_KEY") != "" {
		cfg.APIKey = os.Getenv("API_KEY")
	}
	if cfg.APIKey == "" {
		return fmt.Errorf("API_KEY is missing")
	}

	if interval := os.Getenv("UPDATE_INTERVAL"); interval != "" {
		if val, err := strconv.Atoi(interval); err == nil {
			cfg.UpdateInterval = val
		}
	}

	if offset := os.Getenv("INIT_OFFSET"); offset != "" {
		if val, err := strconv.Atoi(offset); err == nil {
			cfg.InitOffset = val
		}
	}

	return nil
}

// checkConfigValues checks if the config values are valid
func checkConfigValues(cfg *Config) error {
	// check if API_KEY is empty
	if cfg.APIKey == "" {
		return fmt.Errorf("API_KEY is empty")
	}

	// check if INIT_OFFSET is greater than 0
	if cfg.InitOffset < 0 {
		return fmt.Errorf("INIT_OFFSET must be greater than or equal 0")
	}

	// check if UPDATE_INTERVAL is greater than 0
	if cfg.UpdateInterval <= 0 {
		return fmt.Errorf("UPDATE_INTERVAL must be greater than 0")
	}

	return nil
}
