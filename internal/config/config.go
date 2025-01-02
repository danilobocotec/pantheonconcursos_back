package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Database DatabaseConfig
	Server   ServerConfig
	JWT      JWTConfig
	CORS     CORSConfig
	OAuth    OAuthConfig
	Admin    AdminConfig
	Asaas    AsaasConfig
}

type OAuthConfig struct {
	GoogleClientID     string
	GoogleClientSecret string
	FacebookAppID      string
	FacebookAppSecret  string
	RedirectURL        string
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
	Timezone string
}

type ServerConfig struct {
	Port   string
	Env    string
	Host   string
	Scheme string
}

type JWTConfig struct {
	Secret     string
	Expiration string
}

type CORSConfig struct {
	Origin string
}

type AdminConfig struct {
	Secret string
}

type AsaasConfig struct {
	BaseURL string
	Token   string
}

func LoadConfig() (*Config, error) {
	// Load .env file
	_ = godotenv.Load()

	cfg := &Config{
		Database: DatabaseConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "5432"),
			User:     getEnv("DB_USER", "postgres"),
			Password: getEnv("DB_PASSWORD", ""),
			DBName:   getEnv("DB_NAME", "thepantheon_db"),
			SSLMode:  getEnv("DB_SSLMODE", "require"),
			Timezone: getEnv("DB_TIMEZONE", "UTC"),
		},
		Server: ServerConfig{
			Port:   getEnv("SERVER_PORT", "8080"),
			Env:    getEnv("SERVER_ENV", "development"),
			Host:   getEnv("SERVER_HOST", "localhost:8080"),
			Scheme: getEnv("SERVER_SCHEME", "http"),
		},
		JWT: JWTConfig{
			Secret:     getEnv("JWT_SECRET", "your_secret_key"),
			Expiration: getEnv("JWT_EXPIRATION", "24h"),
		},
		CORS: CORSConfig{
			Origin: getEnv("CORS_ORIGIN", "http://localhost:3000"),
		},
		OAuth: OAuthConfig{
			GoogleClientID:     getEnv("GOOGLE_CLIENT_ID", ""),
			GoogleClientSecret: getEnv("GOOGLE_CLIENT_SECRET", ""),
			FacebookAppID:      getEnv("FACEBOOK_APP_ID", ""),
			FacebookAppSecret:  getEnv("FACEBOOK_APP_SECRET", ""),
			RedirectURL:        getEnv("OAUTH_REDIRECT_URL", "http://localhost:8080/api/v1"),
		},
		Admin: AdminConfig{
			Secret: getEnv("ADMIN_SECRET", ""),
		},
		Asaas: AsaasConfig{
			BaseURL: getEnv("ASAAS_BASE_URL", "https://api-sandbox.asaas.com/"),
			Token:   getEnv("ASAAS_TOKEN", ""),
		},
	}

	return cfg, nil
}

func (c *DatabaseConfig) GetDSN() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=%s",
		c.Host,
		c.Port,
		c.User,
		c.Password,
		c.DBName,
		c.SSLMode,
		c.Timezone,
	)
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
