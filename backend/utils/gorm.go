package utils

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetGormConnection() (*gorm.DB, error) {
	connStr := "user=postgres password=postgres1 dbname=JobTrackerDB sslmode=verify-full"
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
		return nil, err
	}

	return db, nil
}
