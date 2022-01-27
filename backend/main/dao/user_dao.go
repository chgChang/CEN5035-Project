package dao

import (
	"Project01/main/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type UserDao interface {
	InsertUser(user models.User)
	UpdateUser(user models.User)
	DeleteUser(user models.User)
	FindAllUser() []models.User
	CloseDB()
}

type DatabaseUser struct {
	connection *gorm.DB
}

func (db *DatabaseUser) InsertUser(user models.User) {
	db.connection.Create(&user)
}

func (db *DatabaseUser) UpdateUser(user models.User) {
	db.connection.Save(&user)
}

func (db *DatabaseUser) DeleteUser(user models.User) {
	db.connection.Delete(&user)
}

func (db *DatabaseUser) FindAllUser() []models.User {
	var userList []models.User
	db.connection.Find(&userList)
	return userList
}

func InitUserDao() UserDao {
	dsn := "root:Zhangchi1@tcp(127.0.0.1:3306)/gin1?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		return nil
	}
	return &DatabaseUser{connection: db}
}

func (db *DatabaseUser) CloseDB() {
	database, _ := db.connection.DB()
	err := database.Close()
	if err != nil {
		panic("Failed to close database")
	}
}
