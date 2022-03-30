package service

import (
	"backend/main/dao"
	"backend/main/form"
	"backend/main/pojo"
	"backend/main/vo"
	"errors"
	"github.com/google/uuid"
	"time"
)

type OrderService interface {
	Checkout(checkoutForm form.CheckOutForm, email string) error
	GetHistory(email string) ([]vo.OrderHistoryVo, error)
	DeleteOrderByEmail(email string) error
}

type orderService struct {
	orderDao     dao.OrderDao
	orderItemDao dao.OrderItemDao
	cartDao      dao.CartDao
	itemDao      dao.ItemDao
}

func (service *orderService) DeleteOrderByEmail(email string) error {

	//Justify if the order of this user is empty
	orderList := service.orderDao.FindOrderByEmail(email)
	if len(orderList) == 0 {
		err := errors.New("order for this user is empty, email = " + email)
		return err
	}

	//Justify if the orderItem of this user is empty
	orderItemList := service.orderItemDao.FindOrderItemByEmail(email)
	if len(orderItemList) == 0 {
		err := errors.New("order for this user is empty, email = " + email)
		return err
	}

	service.orderDao.DeleteOrderByEmail(email)
	service.orderItemDao.DeleteOrderItemByEmail(email)
	return nil

}

func (service *orderService) GetHistory(email string) ([]vo.OrderHistoryVo, error) {
	orderList := service.orderDao.FindOrderByEmail(email)
	orderItemList := service.orderItemDao.FindOrderItemByEmail(email)

	if len(orderList) == 0 && len(orderItemList) == 0 {
		err := errors.New("order history is empty")
		return nil, err
	}

	var orderHistoryList []vo.OrderHistoryVo
	for i := 0; i < len(orderList); i++ {
		var orderHistory vo.OrderHistoryVo
		var itemVoList []vo.ItemVo
		var totalPrice float64 = 0
		orderHistory.OrderId = orderList[i].OrderNo
		orderHistory.OrderDate = orderList[i].CreateTime
		for j := 0; j < len(orderItemList); j++ {
			var itemVo vo.ItemVo
			if orderList[i].OrderNo == orderItemList[j].OrderNo {
				itemVo.ItemId = orderItemList[j].ItemId
				itemVo.ItemName = orderItemList[j].ItemName
				itemVo.PicUrl = orderItemList[j].ItemImage
				itemVo.Description = orderItemList[j].ItemDescription
				itemVo.Quantity = orderItemList[j].Quantity
				itemVo.Price = orderItemList[j].UnitPrice
				itemVoList = append(itemVoList, itemVo)
				totalPrice = totalPrice + orderItemList[j].TotalPrice
			}
		}
		orderHistory.ItemVoList = itemVoList
		orderHistory.TotalPrice = totalPrice
		orderHistoryList = append(orderHistoryList, orderHistory)
	}
	return orderHistoryList, nil
}

func (service *orderService) Checkout(checkoutForm form.CheckOutForm, email string) error {
	cartItemList := service.cartDao.FindCartByEmail(email)

	//cart must not be empty
	if len(cartItemList) == 0 {
		err := errors.New("cart is empty")
		return err
	}

	var order pojo.Order
	var orderItemList []pojo.OrderItem

	//generate an order number randomly
	orderNo := uuid.New().String()

	order.OrderNo = orderNo
	order.Email = email
	order.Name = checkoutForm.Name
	order.Address = checkoutForm.Address
	order.Phone = checkoutForm.Phone
	order.CreateTime = time.Now()
	service.orderDao.InsertOrder(order)

	var itemIdList []int
	for i := 0; i < len(cartItemList); i++ {
		itemIdList = append(itemIdList, cartItemList[i].ItemId)
	}
	itemList := service.itemDao.FindItemByIdList(itemIdList)

	for i := 0; i < len(cartItemList); i++ {
		var orderItem pojo.OrderItem

		orderItem.ItemId = cartItemList[i].ItemId
		orderItem.OrderNo = orderNo
		orderItem.Email = email
		orderItem.Quantity = cartItemList[i].Quantity

		for j := 0; j < len(itemList); j++ {
			if itemList[j].Id == cartItemList[i].ItemId {
				orderItem.ItemName = itemList[j].Name
				orderItem.ItemImage = itemList[j].PicUrl
				orderItem.UnitPrice = itemList[j].Price
				orderItem.ItemDescription = itemList[j].Description
			}
		}
		orderItem.TotalPrice = orderItem.UnitPrice * float64(orderItem.Quantity)
		orderItemList = append(orderItemList, orderItem)
	}
	service.orderItemDao.InsertOrderItemList(orderItemList)

	//after place the order, remove all items in the cart
	service.cartDao.DeleteCartByEmail(email)
	return nil
}

func NewOrderService(orderDao dao.OrderDao, orderItemDao dao.OrderItemDao, cartDao dao.CartDao, itemDao dao.ItemDao) OrderService {
	return &orderService{
		orderDao:     orderDao,
		orderItemDao: orderItemDao,
		cartDao:      cartDao,
		itemDao:      itemDao,
	}
}
