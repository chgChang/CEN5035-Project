package dao

import (
	"backend/main/pojo"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type OrderItemDao interface {
	InsertOrderItem(orderItem pojo.OrderItem)
	InsertOrderItemList(orderItemList []pojo.OrderItem)
	FindOrderItemByEmail(email string) []pojo.OrderItem
}

func (db *Database) InsertOrderItem(orderItem pojo.OrderItem) {
	db.connection.Create(&orderItem)
}

func (db *Database) InsertOrderItemList(orderItemList []pojo.OrderItem) {
	db.connection.Create(&orderItemList)
}

func (db *Database) FindOrderItemByEmail(email string) []pojo.OrderItem {
	var orderItemList []pojo.OrderItem
	db.connection.Where("email = ?", email).Find(&orderItemList)
	return orderItemList
}

func InitOrderItemDao() OrderItemDao {
	dsn := "root:software@tcp(127.0.0.1:3306)/amazon?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}
	err = db.AutoMigrate(&pojo.OrderItem{})
	if err != nil {
		return nil
	}
	return &Database{connection: db}
}
