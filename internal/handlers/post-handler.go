package handlers

import (
	"blogapi/internal/dtos"
	"blogapi/internal/services"
	"blogapi/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

// PostHandler định nghĩa các phương thức xử lý bài viết
type PostHandler interface {
	GetAllPosts(context *gin.Context)
	GetPostById(context *gin.Context)
	CreatePost(context *gin.Context)
	DeletePost(context *gin.Context)
}

type postHandler struct {
	postService services.PostService
	logger      zerolog.Logger
}

func NewPostHandler(service services.PostService, logger zerolog.Logger) *postHandler {
	return &postHandler{
		postService: service,
		logger:      logger,
	}
}

func (handler *postHandler) GetAllPosts(context *gin.Context) {
	offset := utils.GetQueryInt(context, "offset", 0)
	limit := utils.GetQueryInt(context, "limit", 10)

	posts := handler.postService.GetAll(offset, limit)
	context.JSON(http.StatusOK, utils.GetResponse(posts))
}

func (handler *postHandler) GetPostById(context *gin.Context) {
	postId := context.Param("id")
	post, err := handler.postService.GetById(postId)
	if err != nil {
		handler.logger.Error().Err(err).Msg("Post not found")
		context.JSON(http.StatusNotFound, utils.GetErrorResponse("Post not found"))
		return
	}

	context.JSON(http.StatusOK, utils.GetResponse(post))
}

func (handler *postHandler) CreatePost(context *gin.Context) {
	var postInput dtos.CreatePostInput
	if err := context.ShouldBindJSON(&postInput); err != nil {
		handler.logger.Error().Err(err).Msg("Invalid post input")
		context.JSON(http.StatusBadRequest, utils.GetErrorResponse(err.Error()))
		return
	}
	userId := utils.GetUserIdFromContext(context)
	post, err := handler.postService.CreatePost(userId, &postInput)
	if err != nil {
		handler.logger.Error().Err(err).Msg("Failed to create post")
		context.JSON(http.StatusInternalServerError, utils.GetErrorResponse("Failed to create post"))
		return
	}
	context.JSON(http.StatusCreated, utils.GetResponse(post))
}

func (handler *postHandler) DeletePost(context *gin.Context) {
	postId := context.Param("id")
	if err := handler.postService.Delete(postId); err != nil {
		handler.logger.Error().Err(err).Msg("Failed to delete post")
		context.JSON(http.StatusInternalServerError, utils.GetErrorResponse("Failed to delete post"))
		return
	}

	context.JSON(http.StatusOK, utils.GetResponse(gin.H{"message": "Post deleted successfully"}))
}
