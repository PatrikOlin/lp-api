package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var GlobalDB *gorm.DB

func InitDatabase() (err error) {
	dsn := "host=db user=lazypants password=lazypants dbname=lazypants port=5432 sslmode=disable"
	GlobalDB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return
	}

	return
}
