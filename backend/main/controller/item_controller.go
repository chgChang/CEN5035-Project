package controller

import (
	"backend/main/pojo"
	"backend/main/service"
	"github.com/gin-gonic/gin"
	"strconv"
)

type ItemController interface {
	GetItemList(c *gin.Context) ([]pojo.Item, error)
	SearchItem(c *gin.Context) ([]pojo.Item, error)
	SearchItemById(c *gin.Context) (pojo.Item, error)
}

type itemController struct {
	itemService service.ItemService
}

func (controller *itemController) GetItemList(c *gin.Context) ([]pojo.Item, error) {
	itemList, err := controller.itemService.GetItemList()
	if err != nil {
		return nil, err
	}
	return itemList, nil
}

func (controller *itemController) SearchItem(c *gin.Context) ([]pojo.Item, error) {
	keyword := c.Query("searchKey")
	itemList, err := controller.itemService.SearchItem(keyword)
	if err != nil {
		return nil, err
	}
	return itemList, nil
}

func (controller *itemController) SearchItemById(c *gin.Context) (pojo.Item, error) {
	idInUrl := c.Query("id")
	id, err := strconv.Atoi(idInUrl)
	if err != nil {
		return pojo.Item{}, err
	}
	item, err := controller.itemService.SearchItemById(id)
	if err != nil {
		return pojo.Item{}, err
	}
	return item, nil
}

func NewItemController(itemService service.ItemService) ItemController {
	return &itemController{
		itemService: itemService,
	}
}
