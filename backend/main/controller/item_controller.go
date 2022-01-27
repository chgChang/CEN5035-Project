package controller

import (
	"Project01/main/models"
	"Project01/main/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"strconv"
)

type ItemController interface {
	InsertItem(c *gin.Context) error
	UpdateItem(c *gin.Context) error
	DeleteItem(c *gin.Context) error
	FindAllItem() []models.Item
}

type itemController struct {
	itemService service.ItemService
}

var validateItem *validator.Validate

func (controller *itemController) InsertItem(c *gin.Context) error {
	var item models.Item
	err := c.ShouldBindJSON(&item)
	if err != nil {
		return err
	}
	err = validateItem.Struct(item)
	if err != nil {
		return err
	}
	controller.itemService.InsertItem(item)
	return nil
}

func (controller *itemController) UpdateItem(c *gin.Context) error {
	var item models.Item
	id, err := strconv.ParseInt(c.Param("id"), 0, 0)
	if err != nil {
		return err
	}
	item.Id = int(id)
	controller.itemService.UpdateItem(item)
	return nil
}

func (controller *itemController) DeleteItem(c *gin.Context) error {
	var item models.Item
	id, err := strconv.ParseInt(c.Param("id"), 0, 0)
	if err != nil {
		return err
	}
	item.Id = int(id)
	controller.itemService.DeleteItem(item)
	return nil
}

func (controller *itemController) FindAllItem() []models.Item {
	return controller.itemService.FindAllItem()
}

func NewItem(itemService service.ItemService) ItemController {
	validateItem = validator.New()
	return &itemController{
		itemService: itemService,
	}
}
