package initializers

import (
	"os"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB


func ConnectToDb() {
	var err error
	var host = os.Getenv("DB_HOST")
	var user = os.Getenv("DB_USER")
	var password = os.Getenv("DB_PASSWORD")
	var dbname = os.Getenv("DB_NAME")
	var port = os.Getenv("DB_PORT")

	dsn := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}
}