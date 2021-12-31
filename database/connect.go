package database

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	models "gofiber_restapi/app/models"
)

var DB *gorm.DB

func env(key string) string {
	return os.Getenv(key)
}

func ConnectDB() {
	var err error

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", env("DB_USER"), env("DB_PASSWORD"), env("DB_HOST"), env("DB_PORT"), env("DB_NAME"))

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("failed to connect database")
	}

	fmt.Println("Connected to database:", env("DB_NAME"))

	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Note{})
	fmt.Println("Database migrated")
}
