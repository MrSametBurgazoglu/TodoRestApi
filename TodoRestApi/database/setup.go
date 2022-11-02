package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"todoRestApi/models"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "TodoUser:76T4376T43@tcp(localhost:3306)/todo_database?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	migrateErr := db.AutoMigrate(&models.Todo{})
	if migrateErr != nil {
		panic("migration error")
		return
	}

	DB = db
}
