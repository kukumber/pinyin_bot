// Package config handles command line flags and environment variables.
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
	Debug          bool
}

// ParseFlags returns a Config built from environment variables and command
// line flags. Flags take precedence over environment variables, which take
// precedence over defaults.
func ParseFlags() (Config, error) {
	cfg, err := fromEnv()
	if err != nil {
		return cfg, err
	}

	flag.StringVar(&cfg.APIKey, "api-key", cfg.APIKey, "Telegram API key")
	flag.IntVar(&cfg.InitOffset, "init-offset", cfg.InitOffset, "Initial Telegram offset")
	flag.IntVar(&cfg.UpdateInterval, "update-interval", cfg.UpdateInterval, "Update interval in seconds")
	flag.BoolVar(&cfg.Debug, "debug", cfg.Debug, "Enable Telegram API debug logging")
	flag.Parse()

	if err := cfg.validate(); err != nil {
		return Config{}, err
	}

	return cfg, nil
}

// fromEnv reads configuration defaults from environment variables
func fromEnv() (Config, error) {
	cfg := Config{
		APIKey:         os.Getenv("API_KEY"),
		InitOffset:     0,
		UpdateInterval: 60,
	}

	if offset := os.Getenv("INIT_OFFSET"); offset != "" {
		val, err := strconv.Atoi(offset)
		if err != nil {
			return cfg, fmt.Errorf("invalid INIT_OFFSET %q: %w", offset, err)
		}
		cfg.InitOffset = val
	}

	if interval := os.Getenv("UPDATE_INTERVAL"); interval != "" {
		val, err := strconv.Atoi(interval)
		if err != nil {
			return cfg, fmt.Errorf("invalid UPDATE_INTERVAL %q: %w", interval, err)
		}
		cfg.UpdateInterval = val
	}

	if debug := os.Getenv("DEBUG"); debug != "" {
		val, err := strconv.ParseBool(debug)
		if err != nil {
			return cfg, fmt.Errorf("invalid DEBUG %q: %w", debug, err)
		}
		cfg.Debug = val
	}

	return cfg, nil
}

// validate checks if the config values are valid
func (c Config) validate() error {
	if c.APIKey == "" {
		return fmt.Errorf("API key is missing: set --api-key or API_KEY")
	}
	if c.InitOffset < 0 {
		return fmt.Errorf("INIT_OFFSET must be greater than or equal 0")
	}
	if c.UpdateInterval <= 0 {
		return fmt.Errorf("UPDATE_INTERVAL must be greater than 0")
	}

	return nil
}
