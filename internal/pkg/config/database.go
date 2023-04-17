package config

import (
	"fmt"
	"gobio/internal/app/entity"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabase(configuration Config) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		configuration.Get("DB_HOST", "localhost"), configuration.Get("DB_USER", "postgres"), configuration.Get("DB_PASSWORD", "postgres"),
		configuration.Get("DB_NAME", "gobio"), configuration.Get("DB_PORT", "5432"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&entity.User{})
	db.AutoMigrate(&entity.Link{})

	return db
}
