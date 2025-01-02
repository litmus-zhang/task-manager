package api

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/litmus-zhang/task-manager/lib/config"
	"github.com/litmus-zhang/task-manager/lib/db"
)

type Server struct {
	store  db.Store
	router *gin.Engine
}

func NewServer() (*Server, error) {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Error loading config: %s", err)
	}
	store := db.NewStore(cfg)
	server := &Server{
		store: store,
	}

	server.setupRouter()

	return server, nil
}
func (server *Server) setupRouter() {
	router := gin.Default()

	api := router.Group("/api/v1")

	api.GET("/health", server.healthCheck)

	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
