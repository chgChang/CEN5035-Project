package dao

import (
	"backend/main/pojo"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type UserDao interface {
	Insert(user pojo.User)
	FindUserByEmail(email string) pojo.User
	CloseDB()
}

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

func (db *Database) Insert(user pojo.User) {
	db.connection.Create(&user)
}

func (db *Database) FindUserByEmail(email string) pojo.User {
	var user pojo.User
	db.connection.Where("email = ?", email).First(&user)
	return user
}

func InitUserDao() UserDao {
	dsn := "root:software@tcp(127.0.0.1:3306)/amazon?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}
	err = db.AutoMigrate(&pojo.User{})
	if err != nil {
		return nil
	}
	return &Database{connection: db}
}
