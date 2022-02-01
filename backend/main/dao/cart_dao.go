package dao

import (
	"backend/main/pojo"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type CartDao interface {
	InsertCart(cart pojo.Cart)
	FindCartByEmailAndItemId(email string, itemId int) pojo.Cart
	UpdateCart(cart pojo.Cart)
	FindCartByEmail(email string) []pojo.Cart
}

func (db *Database) FindCartByEmail(email string) []pojo.Cart {
	var cartList []pojo.Cart
	db.connection.Where("email = ?", email).Find(&cartList)
	return cartList
}

func (db *Database) UpdateCart(cart pojo.Cart) {
	db.connection.Save(cart)
}

func (db *Database) InsertCart(cart pojo.Cart) {
	db.connection.Create(&cart)
}

func (db *Database) FindCartByEmailAndItemId(email string, itemId int) pojo.Cart {
	var cart pojo.Cart
	db.connection.Where("email = ? AND item_id = ?", email, itemId).First(&cart)
	return cart
}

func InitCartDao() CartDao {
	dsn := "root:software@tcp(127.0.0.1:3306)/amazon?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}
	err = db.AutoMigrate(&pojo.Cart{})
	if err != nil {
		return nil
	}
	return &Database{connection: db}
}
