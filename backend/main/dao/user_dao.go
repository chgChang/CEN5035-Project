package dao

import (
	"backend/main/pojo"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type UserDao interface {
	InsertUser(user pojo.User)
	FindUserByEmail(email string) pojo.User
	DeleteUserById(id int)
	FindUserById(id int) pojo.User
	CloseDB()
}

func (db *Database) FindUserById(id int) pojo.User {
	var user pojo.User
	db.connection.Where("id = ?", id).First(&user)
	return user
}

func (db *Database) DeleteUserById(id int) {
	db.connection.Where("id = ?", id).Delete(pojo.User{})
}

func (db *Database) InsertUser(user pojo.User) {
	db.connection.Create(&user)
}

func (db *Database) FindUserByEmail(email string) pojo.User {
	var user pojo.User
	db.connection.Where("email = ?", email).First(&user)
	return user
}

func InitUserDao() UserDao {
	dsn := "root:softwareengineering@tcp(34.73.130.106:3306)/amazon?charset=utf8mb4&parseTime=True&loc=Local"
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
