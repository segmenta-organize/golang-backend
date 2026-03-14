package configs

import (
	"fmt"
	"log"

	"segmenta/src/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Database *gorm.DB

func ConnectDatabase(currentConfig *models.AppConfig) error {
	dsn := currentConfig.DBConnection
	if dsn == "" {
		sslMode := currentConfig.DBSSLMode
		if sslMode == "" {
			sslMode = "disable"
		}
		dsn = fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
			currentConfig.DBHost,
			currentConfig.DBPort,
			currentConfig.DBUser,
			currentConfig.DBPassword,
			currentConfig.DBName,
			sslMode,
		)
	}

	db, errorHandler := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if errorHandler != nil {
		return fmt.Errorf("[DB] Failed to connect to database: %w", errorHandler)
	}

	// Migrate each model individually so one failure doesn't block the rest
	modelsToMigrate := []interface{}{
		&models.User{},
		&models.Course{},
		&models.Chapter{},
		&models.ExploreCourse{},
		&models.ExploreChapter{},
		&models.Category{},
	}

	for _, model := range modelsToMigrate {
		if err := db.AutoMigrate(model); err != nil {
			log.Printf("[DB] Warning: auto migrate issue for %T: %v", model, err)
		}
	}

	Database = db
	return nil
}
