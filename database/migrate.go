package database

import (
	"github.com/akposieyefa/crud-gorilla/models"
)

func init() {
	ConectToDB()
}

func MigrateDatabaseTables() {
	DB.AutoMigrate(&models.User{})
}
