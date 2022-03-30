package dao

import (
	"backend/main/pojo"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type OrderDao interface {
	InsertOrder(order pojo.Order)
	FindOrderByEmail(email string) []pojo.Order
	DeleteOrderByEmail(email string)
}

func (db *Database) DeleteOrderByEmail(email string) {
	db.connection.Where("email = ?", email).Delete(pojo.Order{})
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
	dsn := "root:123456@tcp(35.226.149.129:3306)/amazon?charset=utf8mb4&parseTime=True&loc=Local"
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
