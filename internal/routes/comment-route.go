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

func CommentRoute(db *gorm.DB, route *gin.RouterGroup, logger *zerolog.Logger) {
	var (
		commentRepository repositories.CommentRepository = repositories.NewCommentRepository(db)
		commentService    services.CommentService        = services.NewCommentService(commentRepository)
		commentHandler    handlers.CommentHandler        = handlers.NewCommentHandler(commentService, logger)
	)

	route.GET("", commentHandler.GetAllComments)
	route.POST("", middlewares.AuthenticateJWT(), commentHandler.CreateComment)
	route.PUT("/:id", middlewares.AuthenticateJWT(), commentHandler.UpdateComment)
	route.DELETE("/delete/:id", middlewares.AuthenticateJWT(), commentHandler.DeleteComment)
}
