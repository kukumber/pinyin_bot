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
	flag.IntVar(&cfg.UpdateInterval, "update-interval", 0, "Update interval in seconds")
	flag.Parse()

	err := parseEnv(&cfg)
	if err != nil {
		return cfg, err
	}

	err = checkConfigValues(&cfg)
	if err != nil {
		return cfg, err
	}

	return cfg, nil
}

// parseEnv parses environment variables and overrides the config values if empty
func parseEnv(cfg *Config) error {
	var err error

	// parse API_KEY
	if cfg.APIKey == "" && os.Getenv("API_KEY") != "" {
		cfg.APIKey = os.Getenv("API_KEY")
	} else if cfg.APIKey == "" {
		return fmt.Errorf("API_KEY is empty")
	}

	// parse UPDATE_INTERVAL
	if cfg.UpdateInterval == 0 && os.Getenv("UPDATE_INTERVAL") != "" {
		cfg.UpdateInterval, err = strconv.Atoi(os.Getenv("UPDATE_INTERVAL"))
		if err != nil {
			return fmt.Errorf("failed to parse UPDATE_INTERVAL: %w", err)
		}
	} else if cfg.UpdateInterval == 0 {
		cfg.UpdateInterval = 60
	}

	// parse INIT_OFFSET
	if cfg.InitOffset == 0 && os.Getenv("INIT_OFFSET") != "" {
		cfg.InitOffset, err = strconv.Atoi(os.Getenv("INIT_OFFSET"))
		if err != nil {
			return fmt.Errorf("failed to parse INIT_OFFSET: %w", err)
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
