package handlers

import (
	"blogapi/internal/dtos"
	"blogapi/internal/services"
	"blogapi/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

type CommentHandler interface {
	GetAllComments(context *gin.Context)
	CreateComment(context *gin.Context)
	UpdateComment(context *gin.Context)
	DeleteComment(context *gin.Context)
}

type commentHandler struct {
	commentService services.CommentService
	logger         zerolog.Logger
}

func NewCommentHandler(service services.CommentService, logger zerolog.Logger) *commentHandler {
	return &commentHandler{
		commentService: service,
		logger:         logger,
	}
}

func (handler *commentHandler) GetAllComments(context *gin.Context) {
	postId := context.Param("postId")
	offset := utils.GetQueryInt(context, "offset", 0)
	limit := utils.GetQueryInt(context, "limit", 10)

	comments := handler.commentService.GetAllComment(postId, offset, limit)
	context.JSON(http.StatusOK, utils.GetResponse(comments))
}

func (handler *commentHandler) CreateComment(context *gin.Context) {
	var commentInput dtos.CommentInput
	if err := context.ShouldBindJSON(&commentInput); err != nil {
		handler.logger.Error().Err(err).Msg("Invalid comment input")
		context.JSON(http.StatusBadRequest, utils.GetErrorResponse(err.Error()))
		return
	}

	userId := utils.GetUserIdFromContext(context)
	comment := handler.commentService.CreateComment(userId, &commentInput)
	context.JSON(http.StatusCreated, utils.GetResponse(comment))
}

func (handler *commentHandler) UpdateComment(context *gin.Context) {
	commentId := context.Param("id")
	var commentInput dtos.CommentInput
	if err := context.ShouldBindJSON(&commentInput); err != nil {
		handler.logger.Error().Err(err).Msg("Invalid update input")
		context.JSON(http.StatusBadRequest, utils.GetErrorResponse(err.Error()))
		return
	}

	userId := utils.GetUserIdFromContext(context)
	updatedComment, err := handler.commentService.UpdateComment(commentId, userId, &commentInput)
	if err != nil {
		handler.logger.Error().Err(err).Msg("Failed to update comment")
		context.JSON(http.StatusInternalServerError, utils.GetErrorResponse(err.Error()))
		return
	}

	context.JSON(http.StatusOK, utils.GetResponse(updatedComment))
}

func (handler *commentHandler) DeleteComment(context *gin.Context) {
	commentId := context.Param("id")
	if err := handler.commentService.Delete(commentId); err != nil {
		handler.logger.Error().Err(err).Msg("Failed to delete comment")
		context.JSON(http.StatusInternalServerError, utils.GetErrorResponse(err.Error()))
		return
	}

	context.JSON(http.StatusOK, utils.GetResponse(gin.H{"message": "Comment deleted successfully"}))
}
