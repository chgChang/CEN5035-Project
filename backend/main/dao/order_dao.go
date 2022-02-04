package dao

import (
	"backend/main/pojo"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type OrderDao interface {
	InsertOrder(order pojo.Order)
	FindOrderByEmail(email string) []pojo.Order
}

func (db *Database) InsertOrder(order pojo.Order) {
	db.connection.Create(&order)
}

func (db *Database) FindOrderByEmail(email string) []pojo.Order {
	var orderList []pojo.Order
	db.connection.Where("email = ?", email).Find(&orderList)
	return orderList
}

func InitOrderDao() OrderDao {
	dsn := "root:softwareengineering@tcp(34.73.22.78:3306)/amazon?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}
	err = db.AutoMigrate(&pojo.Order{})
	if err != nil {
		return nil
	}
	return &Database{connection: db}
}
