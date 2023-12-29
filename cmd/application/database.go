package application

import (
	"coffee-delivery/main/internal/domain/models"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbInstance struct {
	Db *gorm.DB
}

var DB DbInstance

func ConnectDatabase() {
	dsn := fmt.Sprintf(
		"host=0.0.0.0 user=master password=123master dbname=coffee_delivery port=5432 sslmode=disable TimeZone=America/Winnipeg",
		// os.Getenv("DB_USER"),
		// os.Getenv("DB_PASSWORD"),
		// os.Getenv("DB_NAME"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("Failed to connect to database!. \n", err)
		os.Exit(1)
	}
	log.Println("connected")
	db.Logger = logger.Default.LogMode(logger.Info)

	db.AutoMigrate(&models.User{}, &models.Address{})

	DB = DbInstance{Db: db}
}
