package api

import (
	"github.com/gin-gonic/gin"
	"github.com/litmus-zhang/task-manager/internal/config"
	"github.com/litmus-zhang/task-manager/internal/db"
	"go.uber.org/zap"
)

type Server struct {
	logger *zap.Logger
	config *config.Config
	store  db.Store
	router *gin.Engine
}

func NewServer(cfg *config.Config, store db.Store, logger *zap.Logger) (*Server, error) {
	server := &Server{
		config: cfg,
		store:  store,
		logger: logger,
	}

	server.setupRouter()
	return server, nil
}
func (server *Server) setupRouter() {
	router := gin.Default()

	api := router.Group("/api/v1")

	api.GET("/health", server.healthCheck)

	auth := api.Group("/auth")
	auth.POST("/register", server.registerUser)

	server.router = router
}

func (server *Server) Start() error {
	server.logger.Info("starting server", zap.String("address", server.config.HttpServerAddress))

	return server.router.Run(server.config.HttpServerAddress)
}

func errResponse(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{
		"error": message,
	})
}
