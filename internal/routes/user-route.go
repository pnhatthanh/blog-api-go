package routes

import (
	"blogapi/internal/handlers"
	"blogapi/internal/middlewares"
	"blogapi/internal/repositories"
	"blogapi/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

func UserRoute(db *gorm.DB, route *gin.RouterGroup, logger *zerolog.Logger) {
	var (
		userRepository repositories.UserRepository = repositories.NewUserRepository(db)
		userService    services.UserService        = services.NewUserService(userRepository)
		userHandler    handlers.UserHandler        = handlers.NewUserHandler(userService)
	)
	route.GET("/:id", userHandler.GetUserById)
	route.PUT("/update", middlewares.AuthenticateJWT(), userHandler.UpdateUser)
	route.DELETE("/delete/:id", middlewares.AuthenticateJWT(), userHandler.DeleteUser)
}
