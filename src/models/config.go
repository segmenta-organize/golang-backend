package models

type AppConfig struct {
	// Server configuration
	BackendPort  string
	FrontendPort string
	DomainName   string

	// Database configuration
	DBHost       string
	DBPort       string
	DBUser       string
	DBPassword   string
	DBName       string
	DBSSLMode    string
	DBConnection string

	// JWT configuration
	JWTSecret string

	// SMTP configuration
	SMTPEmail   	   string
	SMTPEmailPassword  string
	SMTPHost           string
	SMTPPort           string
}

type DBConfig struct {
    DBHost     string
    DBPort     string
    DBUser     string
    DBPassword string
    DBName     string
    DBSSLMode  string
	DBConnection string
}
