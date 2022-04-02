package controller

import (
	"backend/main/pojo"
	"backend/main/service"
	"errors"
	"github.com/gin-gonic/gin"
	"strconv"
)

type ItemController interface {
	GetItemList(c *gin.Context) ([]pojo.Item, error)
	SearchItem(c *gin.Context) ([]pojo.Item, error)
	SearchItemById(c *gin.Context) (pojo.Item, error)
	AddItem(c *gin.Context) error
	UpdateItem(c *gin.Context) error
	DeleteItem(c *gin.Context) error
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

//TODO
func (controller *itemController) AddItem(c *gin.Context) error {
	cookie, err := c.Request.Cookie("currentUser")
	if err != nil {
		err = errors.New("user not logged in")
		return err
	}
	//Check if having admin authorization
	email := cookie.Value
	if email != "admin" {
		err = errors.New("you are not authorized")
		return err
	}

	var itemToAdd pojo.Item
	err = c.ShouldBindJSON(&itemToAdd)
	if err != nil {
		return err
	}
	err = validate.Struct(&itemToAdd)
	if err != nil {
		return err
	}
	controller.itemService.AddItem(itemToAdd)
	return nil
}

func (controller *itemController) UpdateItem(c *gin.Context) error {
	cookie, err := c.Request.Cookie("currentUser")
	if err != nil {
		err = errors.New("user not logged in")
		return err
	}
	//Check if having admin authorization
	email := cookie.Value
	if email != "admin" {
		err = errors.New("you are not authorized")
		return err
	}

	var itemToUpdate pojo.Item
	err = c.ShouldBindJSON(&itemToUpdate)
	if err != nil {
		return err
	}
	err = validate.Struct(&itemToUpdate)
	if err != nil {
		return err
	}
	err = controller.itemService.UpdateItem(itemToUpdate)
	if err != nil {
		return err
	}
	return nil
}

func (controller *itemController) DeleteItem(c *gin.Context) error {
	cookie, err := c.Request.Cookie("currentUser")
	if err != nil {
		err = errors.New("user not logged in")
		return err
	}
	//Check if having admin authorization
	email := cookie.Value
	if email != "admin" {
		err = errors.New("you are not authorized")
		return err
	}

	var itemToUpdate pojo.Item
	err = c.ShouldBindJSON(&itemToUpdate)
	if err != nil {
		return err
	}
	err = validate.Struct(&itemToUpdate)
	if err != nil {
		return err
	}

	itemId := itemToUpdate.Id
	err = controller.itemService.DeleteItem(itemId)
	if err != nil {
		return err
	}

	return nil
}

func NewItemController(itemService service.ItemService) ItemController {
	return &itemController{
		itemService: itemService,
	}
}
