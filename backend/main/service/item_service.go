package service

import (
	"backend/main/dao"
	"backend/main/models"
)

type ItemService interface {
	InsertItem(item models.Item) error
	UpdateItem(item models.Item) error
	DeleteItem(item models.Item) error
	FindAllItem() []models.Item
	FindItemById(item models.Item) models.Item
	FindItemByName(item models.Item) []models.Item
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

func (service *itemService) FindItemByName(item models.Item) []models.Item {
	name := item.Name
	return service.itemDao.FindItemByName(name)
}

func (service *itemService) FindItemById(item models.Item) models.Item {
	id := item.Id
	return service.itemDao.FindItemById(id)
}

func NewItem(dao dao.ItemDao) ItemService {
	return &itemService{
		itemDao: dao,
	}
}
