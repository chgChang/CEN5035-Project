package service

import (
	"backend/main/dao"
	"backend/main/pojo"
	"errors"
)

type ItemService interface {
	GetItemList() ([]pojo.Item, error)
	SearchItem(keyword string) ([]pojo.Item, error)
	SearchItemById(id int) (pojo.Item, error)
}

type itemService struct {
	itemDao dao.ItemDao
}

func (service *itemService) GetItemList() ([]pojo.Item, error) {
	itemList := service.itemDao.FindAllItems()
	if len(itemList) == 0 {
		err := errors.New("item list is empty")
		return nil, err
	}
	return itemList, nil
}

func (service *itemService) SearchItem(keyword string) ([]pojo.Item, error) {
	itemList := service.itemDao.FindItemByKeyword(keyword)
	if len(itemList) == 0 {
		err := errors.New("the search result is empty")
		return nil, err
	}
	return itemList, nil
}

func (service *itemService) SearchItemById(id int) (pojo.Item, error) {
	item := service.itemDao.FindItemById(id)
	if item == (pojo.Item{}) {
		err := errors.New("id doesn't exist")
		return pojo.Item{}, err
	}
	return item, nil
}

//TODO
func (service *itemService) AddItem(item pojo.Item) {
	service.itemDao.CreateItem(item)
}

func (service *itemService) UpdateItem(item pojo.Item) error {
	itemId := item.Id
	originItem := service.itemDao.FindItemById(itemId)
	if originItem == (pojo.Item{}) {
		err := errors.New("id doesn't exist")
		return err
	}
	service.itemDao.UpdateItem(item)
	return nil
}

func (service *itemService) DeleteItem(itemId int) error {
	originItem := service.itemDao.FindItemById(itemId)
	if originItem == (pojo.Item{}) {
		err := errors.New("item doesn't exist")
		return err
	}
	service.itemDao.DeleteItem(itemId)
	return nil
}

func NewItemService(itemDao dao.ItemDao) ItemService {
	return &itemService{
		itemDao: itemDao,
	}
}
