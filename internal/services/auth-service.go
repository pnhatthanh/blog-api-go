package services

import (
	"blogapi/internal/dtos"
	"blogapi/internal/repositories"
	"blogapi/pkg/utils"
	"errors"
)

type AuthService interface {
	Login(user *dtos.UserLogin) (string, string, error)
}

type authService struct {
	userRepo repositories.UserRepository
}

func NewAuthService(repo repositories.UserRepository) *authService {
	return &authService{userRepo: repo}
}

func (auth *authService) Login(user *dtos.UserLogin) (string, string, error) {
	_user, err := auth.userRepo.GetByEmail(user.Email)
	if err != nil {
		return "", "", errors.New("invalid email or password")
	}
	if !utils.CheckPasswordHash(_user.Password, user.Password) {
		return "", "", errors.New("invalid email or password")
	}

	accessToken, refreshToken, err := utils.GenerateToken(_user.Id)
	if err != nil {
		return "", "", errors.New("failed to generate token")
	}

	return accessToken, refreshToken, nil
}
