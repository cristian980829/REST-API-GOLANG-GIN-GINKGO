package models

import (
	"gorm.io/driver/sqlite"
   	"gorm.io/gorm"
	"fmt"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(sqlite.Open("./store.db"), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}else{
		fmt.Println("Database connected!")
	}

	database.AutoMigrate(&Product{})

	DB = database
}
