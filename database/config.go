package db

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/misbahabroruddin/task-5-pbi-btpns-misbah-abroruddin/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error load .env file")
		}
		dsn := os.Getenv("DB_URL")

    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

    if err != nil {
        log.Fatalln(err)
    }

    db.AutoMigrate(
			&models.User{},
			&models.Photo{},
		)

    return db
}

