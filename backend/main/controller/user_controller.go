package controller

import (
	"backend/main/pojo"
	"backend/main/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UserController interface {
	Register(c *gin.Context) error
	Login(c *gin.Context) error
	Logout(c *gin.Context) error
}

type userController struct {
	userService service.UserService
}

func (controller *userController) Logout(c *gin.Context) error {
	var user pojo.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		return err
	}
	err = validate.Struct(user)
	if err != nil {
		return err
	}
	err = controller.userService.Logout(user)
	if err != nil {
		return err
	}
	return nil
}

func (controller *userController) Register(c *gin.Context) error {
	var user pojo.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		return err
	}
	err = validate.Struct(user)
	if err != nil {
		return err
	}
	err = controller.userService.Register(user)
	if err != nil {
		return err
	}
	return nil
}

func (controller *userController) Login(c *gin.Context) error {
	var user pojo.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		return err
	}
	err = validate.Struct(user)
	if err != nil {
		return err
	}
	err = controller.userService.Login(user)
	if err != nil {
		return err
	}
	return nil
}

var validate *validator.Validate

func New(userService service.UserService) UserController {
	validate = validator.New()
	return &userController{
		userService: userService,
	}
}
