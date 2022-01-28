package controller

import (
	"backend/main/models"
	"backend/main/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"strconv"
)

type UserController interface {
	InsertUser(c *gin.Context) error
	UpdateUser(c *gin.Context) error
	DeleteUser(c *gin.Context) error
	FindAllUser() []models.User
}

type userController struct {
	userService service.UserService
}

var validateUser *validator.Validate

func (controller *userController) InsertUser(c *gin.Context) error {
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		return err
	}
	err = validateUser.Struct(user)
	if err != nil {
		return err
	}
	controller.userService.InsertUser(user)
	return nil
}

func (controller *userController) UpdateUser(c *gin.Context) error {
	var user models.User
	id, err := strconv.ParseInt(c.Param("id"), 0, 0)
	if err != nil {
		return err
	}
	user.Id = int(id)
	controller.userService.UpdateUser(user)
	return nil
}

func (controller *userController) DeleteUser(c *gin.Context) error {
	var user models.User
	id, err := strconv.ParseInt(c.Param("id"), 0, 0)
	if err != nil {
		return err
	}
	user.Id = int(id)
	controller.userService.DeleteUser(user)
	return nil
}

func (controller *userController) FindAllUser() []models.User {
	return controller.userService.FindAllUser()
}

func NewUser(userService service.UserService) UserController {
	validateUser = validator.New()
	return &userController{
		userService: userService,
	}
}
