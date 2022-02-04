package controller

import (
	"backend/main/pojo"
	"backend/main/service"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"time"
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
	cookie, err := c.Request.Cookie("currentUser")
	if err != nil || cookie.Value != user.Email {
		return errors.New("user not logged in")
	}
	//if cookie.Value != user.Email {
	//	err = errors.NewUserService("user not logged in")
	//	return err
	//}
	err = controller.userService.Logout(user)
	if err != nil {
		return err
	}

	expiration := time.Now()
	expiration = expiration.AddDate(0, 0, -1)
	cookieNew := http.Cookie{Name: "currentUser", Value: "", Expires: expiration}
	http.SetCookie(c.Writer, &cookieNew)
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

	expiration := time.Now()
	expiration = expiration.AddDate(0, 0, -1)
	cookieNew := http.Cookie{Name: "currentUser", Value: "", Expires: expiration}
	http.SetCookie(c.Writer, &cookieNew)

	expiration = time.Now()
	expiration = expiration.AddDate(0, 0, 1)
	cookie := http.Cookie{Name: "currentUser", Value: user.Email, Expires: expiration}
	http.SetCookie(c.Writer, &cookie)
	return nil
}

func NewUserController(userService service.UserService) UserController {
	validate = validator.New()
	return &userController{
		userService: userService,
	}
}
