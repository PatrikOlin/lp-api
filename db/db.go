package db

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var GlobalDB *gorm.DB

func InitDatabase() (err error) {
	err = godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	dsn := os.ExpandEnv("host=${POSTGRES_HOST} user=${POSTGRES_USER} dbname=${POSTGRES_DBNAME}  password=${POSTGRES_PASSWORD} sslmode=disable port=5432")
	GlobalDB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("Failed to connect to database")
		return
	}

	return
}
