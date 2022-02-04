package controller

import (
	"backend/main/form"
	"backend/main/service"
	"backend/main/vo"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type OrderController interface {
	Checkout(c *gin.Context) error
	GetHistory(c *gin.Context) ([]vo.OrderHistoryVo, error)
}

type orderController struct {
	orderService service.OrderService
}

func (controller *orderController) GetHistory(c *gin.Context) ([]vo.OrderHistoryVo, error) {
	cookie, err := c.Request.Cookie("currentUser")
	if err != nil {
		err = errors.New("please login first")
		return nil, err
	}
	email := cookie.Value

	orderHistoryVoList, err := controller.orderService.GetHistory(email)
	if err != nil {
		return nil, err
	}
	return orderHistoryVoList, nil
}

func (controller *orderController) Checkout(c *gin.Context) error {
	cookie, err := c.Request.Cookie("currentUser")
	if err != nil {
		err = errors.New("please login first")
		return err
	}
	email := cookie.Value

	var checkoutForm form.CheckOutForm
	err = c.ShouldBindJSON(&checkoutForm)
	if err != nil {
		return err
	}
	err = validate.Struct(&checkoutForm)
	if err != nil {
		return err
	}

	err = controller.orderService.Checkout(checkoutForm, email)
	if err != nil {
		return err
	}
	return nil
}

func NewOrderController(orderService service.OrderService) OrderController {
	validate = validator.New()
	return &orderController{
		orderService: orderService,
	}
}
