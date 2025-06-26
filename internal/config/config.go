package config

import "os"

type Config struct {
	ServerPort    string
	HMACSecret    string
	BootstrapFile string
	PgHost        string
	PgPort        string
	PgUser        string
	PgPassword    string
	PgDatabase    string
	PgSSLMode     string
}

func Load() Config {
	return Config{
		ServerPort:    getEnv("SERVER_PORT", "8080"),
		HMACSecret:    getEnv("HMAC_SECRET", "secret"),
		BootstrapFile: getEnv("BOOTSTRAP_FILE", "./secrets/boostrap-token.txt"),
		PgHost:        getEnv("PG_HOST", "localhost"),
		PgPort:        getEnv("PG_PORT", "5432"),
		PgUser:        getEnv("PG_USER", "postgres"),
		PgPassword:    getEnv("PG_PASSWORD", "postgres"),
		PgDatabase:    getEnv("PG_DATABASE", "itchaca_db"),
		PgSSLMode:     getEnv("PG_SSLMODE", "disable"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
