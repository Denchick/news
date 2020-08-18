package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func Connect() (*gorm.DB, error) {
	dsn := "host=rogue.db.elephantsql.com port=5432 user=acumlegw dbname=acumlegw password=2ZbAnECs3Lfdc9k9sbpRWzcCOHd1mo97"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
		return nil, err
	}
	return db, nil
}
