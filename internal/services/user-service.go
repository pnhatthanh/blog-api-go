package services

import (
	"blogapi/internal/dtos"
	"blogapi/internal/models"
	"blogapi/internal/repositories"
	"blogapi/pkg/utils"
	"errors"
)

type UserService interface {
	CreateUser(user *dtos.CreateUserInput) (*models.User, error)
	UpdateUser(id string, user *dtos.UpdateUserInput) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	GetUserById(id string) (*models.User, error)
	DeleteUser(id string) error
}

type userService struct {
	userRepo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) *userService {
	return &userService{userRepo: repo}
}

func (serv *userService) CreateUser(user *dtos.CreateUserInput) (*models.User, error) {
	existingUser, _ := serv.userRepo.GetByEmail(user.Email)
	if existingUser != nil {
		return nil, errors.New("Email already exists")
	}
	hashedPassword,err:=utils.HashPassword(user.Password)
	if err!=nil{
		return nil, errors.New("server error")
	}
	_user := models.User{
		FullName:     user.FullName,
		Email:    user.Email,
		Password: hashedPassword, 
	}
	serv.userRepo.Create(&_user)
	return &_user, nil
}

func (serv *userService) UpdateUser(id string, user *dtos.UpdateUserInput) (*models.User, error) {
	_user, err := serv.userRepo.GetById(id)
	if err != nil {
		return nil, err
	}
	_user.FullName = user.FullName
	_user.Email = user.Email
	_user.Password = user.Password 

	updatedUser, err := serv.userRepo.Update(_user)
	if err != nil {
		return nil, err
	}
	return updatedUser, nil
}

func (serv *userService) GetUserByEmail(email string) (*models.User, error) {
	return serv.userRepo.GetByEmail(email)
}

func (serv *userService) GetUserById(id string) (*models.User, error) {
	return serv.userRepo.GetById(id)
}

func (serv *userService) DeleteUser(id string) error {
	return serv.userRepo.Delete(id)
}
