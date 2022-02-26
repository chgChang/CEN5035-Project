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
	cookieNew := http.Cookie{Name: "currentUserName", Value: "", Expires: expiration}
	http.SetCookie(c.Writer, &cookieNew)

	expiration = time.Now()
	expiration = expiration.AddDate(0, 0, -1)
	cookieNew2 := http.Cookie{Name: "currentUserEmail", Value: "", Expires: expiration}
	http.SetCookie(c.Writer, &cookieNew2)
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

	var currUser pojo.User
	currUser, err = controller.userService.Login(user)
	if err != nil {
		return err
	}
	//fmt.Println(currUser.Username)
	expiration := time.Now()
	expiration = expiration.AddDate(0, 0, -1)
	cookieNew := http.Cookie{Name: "currentUserEmail", Value: "", Expires: expiration}
	http.SetCookie(c.Writer, &cookieNew)

	expiration = time.Now()
	expiration = expiration.AddDate(0, 0, -1)
	cookieNew2 := http.Cookie{Name: "currentUserEmail", Value: "", Expires: expiration}
	http.SetCookie(c.Writer, &cookieNew2)

	expiration = time.Now()
	expiration = expiration.AddDate(0, 0, 1)
	cookie := http.Cookie{Name: "currentUserEmail", Value: currUser.Email, Expires: expiration}
	http.SetCookie(c.Writer, &cookie)
	
	expiration = time.Now()
	expiration = expiration.AddDate(0, 0, 1)
	cookie2 := http.Cookie{Name: "currentUserName", Value: currUser.Username, Expires: expiration}
	http.SetCookie(c.Writer, &cookie2)
	return nil
}

func NewUserController(userService service.UserService) UserController {
	validate = validator.New()
	return &userController{
		userService: userService,
	}
}
