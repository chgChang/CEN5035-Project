package service

import (
	"backend/main/dao"
	"backend/main/form"
	"backend/main/pojo"
	"backend/main/vo"
	"errors"
)

type CartService interface {
	AddToCart(cartAdd form.CartAddForm, email string) error
	GetCartList(email string) vo.CartVo
}

type cartService struct {
	itemDao dao.ItemDao
	cartDao dao.CartDao
}

func (service *cartService) GetCartList(email string) vo.CartVo {
	cartList := service.cartDao.FindCartByEmail(email)
	itemList := service.itemDao.FindAllItem()
	var cartVo vo.CartVo
	if len(cartList) == 0 {
		var itemVoList []vo.ItemVo
		return vo.CartVo{
			ItemList:   itemVoList,
			TotalPrice: 0,
		}
	}
	var itemVoList []vo.ItemVo
	var totalPrice float64
	for i := 0; i < len(cartList); i++ {
		itemId := cartList[i].ItemId
		quantity := cartList[i].Quantity
		var itemVo vo.ItemVo
		itemVo.ItemId = itemId
		itemVo.Quantity = quantity
		for j := 0; j < len(itemList); j++ {
			if itemId == itemList[j].Id {
				itemVo.ItemName = itemList[j].Name
				itemVo.Price = itemList[j].Price
				itemVo.PicUrl = itemList[j].PicUrl
				itemVo.Description = itemList[j].Description
			}
		}
		itemVoList = append(itemVoList, itemVo)
		totalPrice = totalPrice + float64(quantity)*itemVo.Price
	}
	cartVo.ItemList = itemVoList
	cartVo.TotalPrice = totalPrice
	return cartVo
}

func (service *cartService) AddToCart(cartAdd form.CartAddForm, email string) error {
	id := cartAdd.ItemId
	quantity := cartAdd.Quantity
	itemInDb := service.itemDao.FindItemById(id)
	if itemInDb == (pojo.Item{}) {
		err := errors.New("item doesn't exist")
		return err
	}
	cart := service.cartDao.FindCartByEmailAndItemId(email, id)
	if cart == (pojo.Cart{}) {
		cart := pojo.Cart{
			Email:    email,
			ItemId:   id,
			Quantity: quantity,
		}
		service.cartDao.InsertCart(cart)
	} else {
		cart.Quantity = cart.Quantity + quantity
		service.cartDao.UpdateCart(cart)
	}
	return nil
}

func NewCartService(itemDao dao.ItemDao, cartDao dao.CartDao) CartService {
	return &cartService{
		itemDao: itemDao,
		cartDao: cartDao,
	}
}
