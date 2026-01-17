package config

import (
	"os"
	"time"
)

type Config struct {
	ServerPort      string
	ShutdownTimeout time.Duration
}

func Load() *Config {
	return &Config{
		ServerPort:      getEnv("SERVER_PORT", "8080"),
		ShutdownTimeout: getDurationEnv("SHUTDOWN_TIMEOUT", 5*time.Second),
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func getDurationEnv(key string, defaultValue time.Duration) time.Duration {
	if value, exists := os.LookupEnv(key); exists {
		if duration, err := time.ParseDuration(value); err == nil {
			return duration
		}
	}
	return defaultValue
}
