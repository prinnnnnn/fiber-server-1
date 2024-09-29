package database

import (
	"context"
	"fiber-server-1/internal/adapter/config"
	"fiber-server-1/internal/core/models"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB(ctx *context.Context, config *config.DB) (*gorm.DB, error) {

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		config.Connection,
		config.User,
		config.Password,
		config.Name,
		config.Port,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println(dsn)
		return nil, err
	}

	// Create Table
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Friendship{})

	return db, nil
}
