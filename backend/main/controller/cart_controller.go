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
	UpdateCart(c *gin.Context) error
	RemoveCart(c *gin.Context) error
	DeleteCartByItemId(c *gin.Context) error
	DeleteCartByEmail(c *gin.Context) error
}

type cartController struct {
	cartService service.CartService
}

func (controller *cartController) DeleteCartByEmail(c *gin.Context) error {
	cookie, err := c.Request.Cookie("currentUser")
	if err != nil {
		err = errors.New("please login first")
		return err
	}

	//Justify the admin authorization
	currentEmail := cookie.Value
	if currentEmail != "admin" {
		err = errors.New("you don't have the authorization to do that")
		return err
	}

	email := c.Query("email")

	err = controller.cartService.RemoveCart(email)
	if err != nil {
		return err
	}
	return nil
}

func (controller *cartController) DeleteCartByItemId(c *gin.Context) error {
	var cartDeleteForm form.CartDeleteForm
	err := c.ShouldBindJSON(&cartDeleteForm)
	if err != nil {
		return err
	}
	err = validate.Struct(&cartDeleteForm)
	if err != nil {
		return err
	}

	cookie, err := c.Request.Cookie("currentUser")
	if err != nil {
		err = errors.New("please login first")
		return err
	}
	email := cookie.Value

	err = controller.cartService.DeleteCartByItemId(email, cartDeleteForm)
	if err != nil {
		return err
	}
	return nil
}

func (controller *cartController) RemoveCart(c *gin.Context) error {
	cookie, err := c.Request.Cookie("currentUser")
	if err != nil {
		err = errors.New("please login first")
		return err
	}
	email := cookie.Value

	err = controller.cartService.RemoveCart(email)
	if err != nil {
		return err
	}
	return nil
}

func (controller *cartController) UpdateCart(c *gin.Context) error {
	var cartUpdateForm form.CartUpdateForm
	err := c.ShouldBindJSON(&cartUpdateForm)
	if err != nil {
		return err
	}
	err = validate.Struct(&cartUpdateForm)
	if err != nil {
		return err
	}

	cookie, err := c.Request.Cookie("currentUser")
	if err != nil {
		err = errors.New("please login first")
		return err
	}
	email := cookie.Value

	err = controller.cartService.UpdateCart(email, cartUpdateForm)
	if err != nil {
		return err
	}
	return nil
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
