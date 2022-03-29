package dao

import (
	"backend/main/pojo"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type CartDao interface {
	InsertCart(cart pojo.Cart)
	FindCartByEmailAndItemId(email string, itemId int) pojo.Cart
	UpdateCart(email string, itemId int, quantity int)
	FindCartByEmail(email string) []pojo.Cart
	DeleteCartByEmail(email string)
	DeleteCartByEmailAndItemId(email string, itemId int)
}

func (db *Database) DeleteCartByEmail(email string) {
	db.connection.Where("email = ?", email).Delete(pojo.Cart{})
}

func (db *Database) FindCartByEmail(email string) []pojo.Cart {
	var cartList []pojo.Cart
	db.connection.Where("email = ?", email).Find(&cartList)
	return cartList
}

func (db *Database) UpdateCart(email string, itemId int, quantity int) {
	db.connection.Model(&pojo.Cart{}).Where("email = ? AND item_id = ?", email, itemId).Update("quantity", quantity)
}

func (db *Database) InsertCart(cart pojo.Cart) {
	db.connection.Create(&cart)
}

func (db *Database) FindCartByEmailAndItemId(email string, itemId int) pojo.Cart {
	var cart pojo.Cart
	db.connection.Where("email = ? AND item_id = ?", email, itemId).First(&cart)
	return cart
}

func (db *Database) DeleteCartByEmailAndItemId(email string, itemId int) {
	db.connection.Where("email = ? AND item_id = ?", email, itemId).Delete(pojo.Cart{})
}

func InitCartDao() CartDao {
	dsn := "root:123456@tcp(35.226.149.129:3306)/amazon?charset=utf8mb4&parseTime=True&loc=Local"
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
