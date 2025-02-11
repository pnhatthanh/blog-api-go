package handlers

import (
	"blogapi/internal/dtos"
	"blogapi/internal/services"
	"blogapi/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	CreateUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	GetUserByEmail(c *gin.Context)
	GetUserById(c *gin.Context)
	DeleteUser(c *gin.Context)
}

type userHandler struct {
	userService services.UserService
}

func NewUserHandler(userService services.UserService) *userHandler {
	return &userHandler{userService: userService}
}

func (h *userHandler) UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var userInput dtos.UpdateUserInput

	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, utils.GetErrorResponse(err.Error()))
		return
	}

	updatedUser, err := h.userService.UpdateUser(id, &userInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.GetErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, utils.GetResponse(updatedUser))
}

func (h *userHandler) GetUserById(c *gin.Context) {
	id := c.Param("id")
	user, err := h.userService.GetUserById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, utils.GetErrorResponse("User not found"))
		return
	}
	c.JSON(http.StatusOK, utils.GetResponse(user))
}

func (h *userHandler) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	if err := h.userService.DeleteUser(id); err != nil {
		c.JSON(http.StatusBadRequest, utils.GetErrorResponse(err.Error()))
		return
	}
	c.JSON(http.StatusOK, utils.GetResponse(gin.H{"message": "User deleted successfully"}))
}
