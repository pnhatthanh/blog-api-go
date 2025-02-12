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

func PostRoute(db *gorm.DB, route *gin.RouterGroup, logger *zerolog.Logger) {
	var (
		postRepository repositories.PostRepository = repositories.NewPostRepository(db)
		postService    services.PostService        = services.NewPostService(postRepository)
		postHandler    handlers.PostHandler        = handlers.NewPostHandler(postService, logger)
	)
	route.GET("", postHandler.GetAllPosts)
	route.GET("/:id", postHandler.GetPostById)
	route.POST("/create", middlewares.AuthenticateJWT(), postHandler.CreatePost)
	route.DELETE("/delete/:id", middlewares.AuthenticateJWT(), postHandler.DeletePost)
}
