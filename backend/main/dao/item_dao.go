package dao

import (
	"Project01/main/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ItemDao interface {
	InsertItem(item models.Item)
	UpdateItem(item models.Item)
	DeleteItem(item models.Item)
	FindAllItem() []models.Item
	CloseDB()
}

type DatabaseItem struct {
	connection *gorm.DB
}

func (db *DatabaseItem) InsertItem(item models.Item) {
	db.connection.Create(&item)
}

func (db *DatabaseItem) UpdateItem(item models.Item) {
	db.connection.Save(&item)
}

func (db *DatabaseItem) DeleteItem(item models.Item) {
	db.connection.Delete(&item)
}

func (db *DatabaseItem) FindAllItem() []models.Item {
	var itemList []models.Item
	db.connection.Find(&itemList)
	return itemList
}

func InitItemDao() ItemDao {
	dsn := "root:Zhangchi1@tcp(127.0.0.1:3306)/gin1?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}
	err = db.AutoMigrate(&models.Item{})
	if err != nil {
		return nil
	}
	return &DatabaseItem{connection: db}
}

func (db *DatabaseItem) CloseDB() {
	database, _ := db.connection.DB()
	err := database.Close()
	if err != nil {
		panic("Failed to close database")
	}
}
