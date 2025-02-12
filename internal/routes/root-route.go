package routes

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

type APIServer struct {
	Address string
	Db      *gorm.DB
}

func NewAPIServer(port string, db *gorm.DB) *APIServer {
	return &APIServer{
		Address: ":" + port,
		Db:      db,
	}
}
func (s *APIServer) Run() error {
	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()
	router := gin.Default()
	route := router.Group("/api/v1")
	authRoute := route.Group("/auth")
	AuthRoute(s.Db, authRoute, &logger)
	postRoute := route.Group("/posts")
	PostRoute(s.Db, postRoute, &logger)
	commentRoute := route.Group("/post/:postId/comments")
	CommentRoute(s.Db, commentRoute, &logger)
	userRoute := route.Group("/user")
	UserRoute(s.Db, userRoute, &logger)
	logger.Info().Msg("Listening on port" + s.Address)
	return router.Run(s.Address)
}
