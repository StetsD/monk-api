package config

import (
	"github.com/stetsd/monk-api/errorsCmd"
	"os"
)

const (
	DbHost  = "DB_HOST"
	DbPort  = "DB_PORT"
	DbName  = "DB_NAME"
	DbUser  = "DB_USER"
	DbPass  = "DB_PASS"
	AppPort = "APP_PORT"
	AppHost = "APP_HOST"
)

var RequiredConfigFields = []string{
	DbHost,
	DbPort,
	DbName,
	DbUser,
	DbPass,
	AppPort,
	AppHost,
}

type Config struct {
	_map map[string]string
}

func EnvParseToConfigMap() (Config, error) {
	config := Config{_map: make(map[string]string)}

	for _, field := range RequiredConfigFields {
		val := os.Getenv(field)

		if val == "" {
			return config, errorsCmd.ErrorCmd("env " + field + " is required")
		}

		config._map[field] = val
	}

	return config, nil
}

func (config *Config) Get(field string) string {
	return config._map[field]
}
