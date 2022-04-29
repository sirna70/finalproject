package database

import (
	"final-project/models"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	// DB is the database connection
	db     *gorm.DB
	err    error
	host   = "localhost"
	user   = "postgres"
	pass   = "admin123"
	dbport = "5432"
	dbname = "finalprogo"
)

func StartDB() {
	config := fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s sslmode=disable", host, user, pass, dbport, dbname)

	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		fmt.Println("Error in connecting to database:", err)
	}

	db.AutoMigrate(&models.User{}, &models.Photo{}, &models.Comment{}, &models.SocialMedia{})
}

func GetDB() *gorm.DB {
	return db
}
