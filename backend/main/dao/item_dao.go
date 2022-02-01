package dao

import (
	"backend/main/pojo"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ItemDao interface {
	FindItemById(id int) pojo.Item
	FindAllItem() []pojo.Item
}

func (db *Database) FindAllItem() []pojo.Item {
	var itemList []pojo.Item
	db.connection.Find(&itemList)
	return itemList
}

func (db *Database) FindItemById(id int) pojo.Item {
	var item pojo.Item
	db.connection.Where("id = ?", id).First(&item)
	return item
}

func InitItemDao() ItemDao {
	dsn := "root:software@tcp(127.0.0.1:3306)/amazon?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}
	err = db.AutoMigrate(&pojo.Item{})
	if err != nil {
		return nil
	}
	return &Database{connection: db}
}
