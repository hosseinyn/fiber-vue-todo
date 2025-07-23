package database

import (
	"backend/models"
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDb() {
	db, err := gorm.Open(sqlite.Open("todo.db?_foreign_keys=true"), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
		return
	}

	db.AutoMigrate(&models.UserModel{})
	db.AutoMigrate(&models.Todo{})

	DB = db
}
