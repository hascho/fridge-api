package database

import (
	"fmt"

	"go-fridge/internal/category"
	"go-fridge/internal/config"
	"go-fridge/internal/item"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init(cfg config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(&category.Category{}, &item.Item{}); err != nil {
		return nil, err
	}

	return db, nil
}
