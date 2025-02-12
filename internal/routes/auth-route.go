package routes

import (
	"blogapi/internal/handlers"
	"blogapi/internal/repositories"
	"blogapi/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

func AuthRoute(db *gorm.DB, route *gin.RouterGroup, logger *zerolog.Logger) {
	var (
		userRepository repositories.UserRepository = repositories.NewUserRepository(db)
		userService    services.UserService        = services.NewUserService(userRepository)
		authService    services.AuthService        = services.NewAuthService(userRepository)
		authHandler    handlers.AuthHandler        = handlers.NewAuthHandler(authService, userService, logger)
	)
	route.POST("/login", authHandler.Login)
	route.POST("/register", authHandler.Register)
	route.POST("/refresh-token", authHandler.RefreshToken)
}
