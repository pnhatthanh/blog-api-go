package handlers

import (
	"blogapi/internal/dtos"
	"blogapi/internal/services"
	"blogapi/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

type AuthHandler interface {
	Login(context *gin.Context)
	Register(context *gin.Context)
	RefreshToken(context *gin.Context)
}

type authHandler struct {
	authService services.AuthService
	logger      zerolog.Logger
	userService services.UserService
}

func NewAuthHandler(authService services.AuthService, userService services.UserService, logger zerolog.Logger) *authHandler {
	return &authHandler{
		authService: authService,
		userService: userService,
		logger:      logger,
	}
}

func (handler *authHandler) Login(context *gin.Context) {
	var userLogin dtos.UserLogin
	if err := context.ShouldBindJSON(&userLogin); err != nil {
		handler.logger.Error().Err(err).Msg("Invalid login request")
		context.JSON(http.StatusBadRequest, utils.GetErrorResponse(err.Error()))
		return
	}

	accessToken, refreshToken, err := handler.authService.Login(&userLogin)
	if err != nil {
		handler.logger.Error().Err(err).Msg("Invalid credentials")
		context.JSON(http.StatusUnauthorized, utils.GetErrorResponse("Invalid email or password"))
		return
	}

	context.JSON(http.StatusOK, utils.GetResponse(gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	}))
}

func (handler *authHandler) Register(context *gin.Context) {
	var createUserInput dtos.CreateUserInput

	if err := context.ShouldBindJSON(&createUserInput); err != nil {
		handler.logger.Error().Err(err).Msg("Invalid registration request")
		context.JSON(http.StatusBadRequest, utils.GetErrorResponse(err.Error()))
		return
	}
	user, err := handler.userService.CreateUser(&createUserInput)
	if err != nil {
		handler.logger.Error().Err(err).Msg("Registration failed")
		context.JSON(http.StatusBadRequest, utils.GetErrorResponse(err.Error()))
		return
	}
	context.JSON(http.StatusOK, utils.GetResponse(user))
}

func (handler *authHandler) RefreshToken(context *gin.Context) {
	var tokenDto dtos.TokenDTO

	if err := context.ShouldBindJSON(&tokenDto); err != nil {
		handler.logger.Error().Err(err).Msg("Invalid refresh token request")
		context.JSON(http.StatusBadRequest, utils.GetErrorResponse(err.Error()))
		return
	}
	userId, err := utils.GetUserIdByToken(tokenDto.AccessToken)
	if err != nil {
		handler.logger.Error().Err(err).Msg("Invalid refresh token")
		context.JSON(http.StatusUnauthorized, utils.GetErrorResponse("Invalid refresh token"))
		return
	}
	accessToken, refreshToken, err := utils.GenerateToken(userId)
	if err != nil {
		handler.logger.Error().Err(err).Msg("Failed to generate new token")
		context.JSON(http.StatusInternalServerError, utils.GetErrorResponse("Failed to generate token"))
		return
	}
	context.JSON(http.StatusOK, utils.GetResponse(gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	}))
}
