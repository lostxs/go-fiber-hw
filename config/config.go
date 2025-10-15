package config

import (
	"log"
	"lostx/go-fiber-hw/pkg/env"
)

type Config struct {
	Logger LoggerConfig
}

type LoggerConfig struct {
	Level  int
	Format string
}

func Load() *Config {
	if err := env.Load(); err != nil {
		log.Printf("[config] warning: %v\n", err)
	}

	cfg := &Config{
		Logger: LoggerConfig{
			Level:  env.Get("LOG_LEVEL", 0),
			Format: env.Get("LOG_FORMAT", "text"),
		},
	}

	return cfg
}
