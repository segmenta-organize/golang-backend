package configs

import (
	"log"
	"os"

	"segmenta/src/models"

	"github.com/joho/godotenv"
)

func LoadAppConfig() *models.AppConfig {
	if errorHandler := godotenv.Load(); errorHandler != nil {
		// Try loading from parent directory (when running from src/)
		if errorHandler2 := godotenv.Load("../.env"); errorHandler2 != nil {
			log.Printf("[ENV] Error loading .env file: %v", errorHandler)
		}
	}

	return &models.AppConfig{
		// Server configuration
		BackendPort: getEnv("BACKEND_PORT", "3344"),
		FrontendPort: getEnv("FRONTEND_PORT", "3000"),
		DomainName:   getEnv("DOMAIN_NAME", "localhost"),
		
		// Database configuration
		DBHost:       os.Getenv("DB_HOST"),
		DBPort:       os.Getenv("DB_PORT"),
		DBUser:       os.Getenv("DB_USER"),
		DBPassword:   os.Getenv("DB_PASSWORD"),
		DBName:       os.Getenv("DB_NAME"),
		DBSSLMode:    os.Getenv("DB_SSL_MODE"),
		DBConnection: os.Getenv("DB_CONNECTION"),

		// JWT configuration
		JWTSecret: os.Getenv("JWT_SECRET"),

		// SMTP configuration
		SMTPEmail:   	   os.Getenv("SMTP_EMAIL"),
		SMTPEmailPassword: os.Getenv("SMTP_EMAIL_PASSWORD"),
		SMTPHost:          os.Getenv("SMTP_HOST"),
		SMTPPort:          os.Getenv("SMTP_PORT"),
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func GetJWTSecretKey() string {
	return getEnv("JWT_SECRET", "default_secret_key")
}