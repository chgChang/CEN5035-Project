package service

import (
	"Project01/main/dao"
	"Project01/main/models"
)

type ItemService interface {
	InsertItem(item models.Item) error
	UpdateItem(item models.Item) error
	DeleteItem(item models.Item) error
	FindAllItem() []models.Item
}

type itemService struct {
	itemDao dao.ItemDao
}

func (service *itemService) InsertItem(item models.Item) error {
	service.itemDao.InsertItem(item)
	return nil
}

func (service *itemService) UpdateItem(item models.Item) error {
	service.itemDao.UpdateItem(item)
	return nil
}

func (service *itemService) DeleteItem(item models.Item) error {
	service.itemDao.DeleteItem(item)
	return nil
}

func (service *itemService) FindAllItem() []models.Item {
	return service.itemDao.FindAllItem()
}

func NewItem(dao dao.ItemDao) ItemService {
	return &itemService{
		itemDao: dao,
	}
}
