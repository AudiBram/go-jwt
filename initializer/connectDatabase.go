package initializer

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	dsn := os.Getenv("DB")
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Cannot connect to database")
	}
}
