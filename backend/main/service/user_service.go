package service

import (
	"backend/main/dao"
	"backend/main/pojo"
	"errors"
	"strconv"
)

type UserService interface {
	Register(user pojo.User) error
	Login(user pojo.User) (pojo.User, error)
	Logout(user pojo.User) error
	DeleteUser(id int) error
}

type userService struct {
	userDao      dao.UserDao
	cartDao      dao.CartDao
	orderDao     dao.OrderDao
	orderItemDao dao.OrderItemDao
}

func (service *userService) DeleteUser(id int) error {

	//Admin user is not deletable.
	if id == 1 {
		err := errors.New("cannot delete the admin user")
		return err
	}

	//Check if the user exists
	userInDB := service.userDao.FindUserById(id)
	if userInDB == (pojo.User{}) {
		err := errors.New("user doesn't exist, userId = " + strconv.Itoa(id))
		return err
	}
	service.userDao.DeleteUserById(id)

	//Delete all the information of this user
	email := userInDB.Email
	//cartInDB := service.cartDao.FindCartByEmail(email)
	//if len(cartInDB) == 0 {
	//	err := errors.New("no items in cart for this user, userId = " + strconv.Itoa(id))
	//	return err
	//}
	service.cartDao.DeleteCartByEmail(email)

	//orderInDB := service.orderDao.FindOrderByEmail(email)
	//if len(orderInDB) == 0 {
	//	err := errors.New("no orders for this user, userId = " + strconv.Itoa(id))
	//	return err
	//}
	service.orderDao.DeleteOrderByEmail(email)

	//orderItemInDB := service.orderItemDao.FindOrderItemByEmail(email)
	//if len(orderItemInDB) == 0 {
	//	err := errors.New("no orders for this user, userId = " + strconv.Itoa(id))
	//	return err
	//}
	service.orderItemDao.DeleteOrderItemByEmail(email)

	return nil
}

func (service *userService) Logout(user pojo.User) error {
	email := user.Email
	userInDB := service.userDao.FindUserByEmail(email)
	if userInDB == (pojo.User{}) {
		err := errors.New("user doesn't exist")
		return err
	} else {
		return nil
	}
}

func (service *userService) Register(user pojo.User) error {
	email := user.Email
	userInDB := service.userDao.FindUserByEmail(email)
	if userInDB != (pojo.User{}) {
		err := errors.New("email already exists")
		return err
	} else {
		service.userDao.InsertUser(user)
		return nil
	}
}

func (service *userService) Login(user pojo.User) (pojo.User, error) {
	email := user.Email
	password := user.Password
	userInDB := service.userDao.FindUserByEmail(email)

	if userInDB == (pojo.User{}) {
		err := errors.New("email or password is wrong")
		return pojo.User{}, err
	} else {
		passwordInDB := userInDB.Password
		if password != passwordInDB {
			err := errors.New("email or password is wrong！")
			return pojo.User{}, err
		} else {
			return userInDB, nil
		}
	}
}

func NewUserService(userDao dao.UserDao, cartDao dao.CartDao, orderDao dao.OrderDao, orderItemDao dao.OrderItemDao) UserService {
	return &userService{
		userDao:      userDao,
		cartDao:      cartDao,
		orderDao:     orderDao,
		orderItemDao: orderItemDao,
	}
}
