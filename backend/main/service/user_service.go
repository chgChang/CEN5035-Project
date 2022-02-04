package service

import (
	"backend/main/dao"
	"backend/main/pojo"
	"errors"
)

type UserService interface {
	Register(user pojo.User) error
	Login(user pojo.User) error
	Logout(user pojo.User) error
}

type userService struct {
	userDao dao.UserDao
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

func (service *userService) Login(user pojo.User) error {
	email := user.Email
	password := user.Password
	userInDB := service.userDao.FindUserByEmail(email)
	if userInDB == (pojo.User{}) {
		err := errors.New("email or password is wrong")
		return err
	} else {
		passwordInDB := userInDB.Password
		if password != passwordInDB {
			err := errors.New("email or password is wrong！")
			return err
		} else {
			return nil
		}
	}
}

func NewUserService(dao dao.UserDao) UserService {
	return &userService{
		userDao: dao,
	}
}
