package database

import (
	"backend/models"
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDb() {
	db, err := gorm.Open(sqlite.Open("todo.db"), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
		return
	}

	db.AutoMigrate(&models.UserModel{}, &models.Todo{})

	DB = db
}
