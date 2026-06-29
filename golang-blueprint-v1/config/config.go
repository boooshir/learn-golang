package config

import (
	"os"
	"strconv"
)

type DBConfig struct {
	DSN         string `mapstructure:"DSN"`
	MaxOpen     int    `mapstructure:"MAX_OPEN"`
	MaxIdle     int    `mapstructure:"MAX_IDLE"`
	MaxLifetime string `mapstructure:"MAX_LIFETIME"`
}

func LoadDBConfig() *DBConfig {
	return &DBConfig{
		DSN:         getEnv("DSN", "postgres://postgres:@localhost:5432/test"),
		MaxOpen:     getEnvAsInt("MAX_OPEN", 20),
		MaxIdle:     getEnvAsInt("MAX_IDLE", 5),
		MaxLifetime: getEnv("MAX_LIFETIME", "5m"),
	}
}

func getEnv(key, defaultValue string) string {
	if value, exist := os.LookupEnv(key); exist {
		return value
	}
	return defaultValue
}
func getEnvAsInt(key string, defaultValue int) int {
	valueStr := getEnv(key, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return defaultValue
}
