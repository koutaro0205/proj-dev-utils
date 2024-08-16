package db

import (
	model "github.com/koutaro0205/backend-template/app/internal/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := "root:password@tcp(db)/backend-templatea-db?charset=utf8mb4&parseTime=true&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	DB = db
	DB.AutoMigrate(&model.User{})
}

func CloseDB() {
	database, err := DB.DB()
	if err != nil {
		panic(err)
	}
	database.Close()
}
