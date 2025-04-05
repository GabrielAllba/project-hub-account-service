package config

import "os"

type Config struct {
	ServerAddr       string
	DBDriver         string
	DBSource         string
	DBName           string
	DefaultDBSource  string
	JWTSecret        string
}

func NewConfig() *Config {
	return &Config{
		ServerAddr:      getEnv("SERVER_ADDR"),
		DefaultDBSource: getEnv("DEFAULT_DB_SOURCE"),
		DBDriver:        getEnv("DB_DRIVER"),
		DBSource:        getEnv("DB_SOURCE"),
		DBName:          getEnv("DB_NAME"),
		JWTSecret:       getEnv("JWT_SECRET"),
	}
}

func getEnv(key string) string {
	return os.Getenv(key)
}
