package service

import (
	"Project01/main/dao"
	"Project01/main/models"
)

type UserService interface {
	InsertUser(user models.User) error
	UpdateUser(user models.User) error
	DeleteUser(user models.User) error
	FindAllUser() []models.User
}

type userService struct {
	userDao dao.UserDao
}

func (service *userService) InsertUser(user models.User) error {
	service.userDao.InsertUser(user)
	return nil
}

func (service *userService) UpdateUser(user models.User) error {
	service.userDao.UpdateUser(user)
	return nil
}

func (service *userService) DeleteUser(user models.User) error {
	service.userDao.DeleteUser(user)
	return nil
}

func (service *userService) FindAllUser() []models.User {
	return service.userDao.FindAllUser()
}

func NewUser(dao dao.UserDao) UserService {
	return &userService{
		userDao: dao,
	}
}
