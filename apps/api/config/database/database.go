package database

import (
	"log"
	"os"

	"github.com/BooBooStory/config/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatal("Gagal koneksi file .env", err)
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	sslmode := os.Getenv("DB_SSLMODE")

	dsn := "host=" + host + " port=" + port + " user=" + user + " password=" + password + " dbname=" + dbname + " sslmode=" + sslmode

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database: " + err.Error())
	}

	err = DB.AutoMigrate(
		&models.User{}, &models.Creator{}, &models.Achievement{},
		&models.Answer{}, &models.Category{}, &models.Progress{},
		&models.Question{}, &models.Quiz{}, &models.ShopItem{},
		&models.Story{}, &models.StoryContent{}, &models.StoryContentRareWord{},
		&models.UserAchievement{}, &models.UserInventory{}, &models.UserQuizResult{},
	)
	if err != nil {
		log.Fatal("failed to migrate database: " + err.Error())
	}
}