package database

import (
	"fmt"
	"time"

	"github.com/At-Sovereign-Technologies/servidor-electoral/internal/models"
	"github.com/At-Sovereign-Technologies/servidor-electoral/internal/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var instance *gorm.DB

func Init() error {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		utils.GetEnv("DB_HOST", "localhost"),
		utils.GetEnv("DB_USER", "postgres"),
		utils.GetEnv("DB_PASSWORD", "postgres"),
		utils.GetEnv("DB_NAME", "postgres"),
		utils.GetEnv("DB_PORT", "5432"),
	)

	var err error
	instance, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent), // disabled logging
	})
	if err != nil {
		return fmt.Errorf("failed to connect database: %w", err)
	}

	db, err := instance.DB()
	if err != nil {
		return fmt.Errorf("failed to get sql.DB from gorm.DB: %w", err)
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(time.Hour)

	return nil
}

func GetDB() *gorm.DB {
	return instance
}

func Migrate() error {
	return instance.AutoMigrate(
		&models.Election{},
		&models.Candidate{},
		&models.ElectionDeployment{},
		&models.VotingPlace{},
		&models.VotingBooth{},
		&models.Voter{},
	)
}
