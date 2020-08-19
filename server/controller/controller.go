package controller

import (
	"github.com/denchick/news/database"
	"gorm.io/gorm"
)

var Db *gorm.DB

func ConnectToDatabase() (err error) {
	Db, err = database.Connect()
	return err
}
