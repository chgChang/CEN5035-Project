package dao

import (
	"gorm.io/gorm"
)

type Database struct {
	connection *gorm.DB
}

func (db *Database) CloseDB() {
	database, _ := db.connection.DB()
	err := database.Close()
	if err != nil {
		panic("Failed to close database")
	}
}
