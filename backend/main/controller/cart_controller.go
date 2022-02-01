package controller

import (
	"backend/main/form"
	"backend/main/service"
	"backend/main/vo"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type CartController interface {
	AddToCart(c *gin.Context) error
	GetCartList(c *gin.Context) (vo.CartVo, error)
}

type cartController struct {
	cartService service.CartService
}

func (controller *cartController) GetCartList(c *gin.Context) (vo.CartVo, error) {
	cookie, err := c.Request.Cookie("currentUser")
	if err != nil {
		err = errors.New("please login first")
		return vo.CartVo{}, err
	}
	email := cookie.Value
	cartVo := controller.cartService.GetCartList(email)
	return cartVo, nil
}

func (controller *cartController) AddToCart(c *gin.Context) error {
	var cartAdd form.CartAddForm
	err := c.ShouldBindJSON(&cartAdd)
	if err != nil {
		return err
	}
	err = validate.Struct(&cartAdd)
	if err != nil {
		return err
	}
	cookie, err := c.Request.Cookie("currentUser")
	if err != nil {
		err = errors.New("please login first")
		return err
	}
	email := cookie.Value
	if email == "" {
		err = errors.New("please login first")
		return err
	}
	err = controller.cartService.AddToCart(cartAdd, email)
	if err != nil {
		return err
	}
	return nil
}

func NewCartController(cartService service.CartService) CartController {
	validate = validator.New()
	return &cartController{
		cartService: cartService,
	}
}
