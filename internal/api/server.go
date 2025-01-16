package api

import (
	"sort"

	"github.com/gin-gonic/gin"
	"github.com/litmus-zhang/momentum-backend/internal/config"
	"github.com/litmus-zhang/momentum-backend/internal/db"
	"github.com/litmus-zhang/momentum-backend/util"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/google"
	"go.uber.org/zap"
)

type ProviderIndex struct {
	Providers    []string
	ProvidersMap map[string]string
}
type Server struct {
	logger     *zap.Logger
	config     *config.Config
	store      db.Store
	router     *gin.Engine
	tokenMaker util.Maker
}

func NewServer(cfg *config.Config, store db.Store, logger *zap.Logger) (*Server, error) {
	tokenMaker, err := util.NewPasetoMaker(cfg.TokenSymmetricKey)
	if err != nil {
		return nil, err
	}
	server := &Server{
		config:     cfg,
		store:      store,
		logger:     logger,
		tokenMaker: tokenMaker,
	}

	server.setupRouter()
	return server, nil
}
func (server *Server) setupRouter() {
	BaseURL := server.config.HttpServerAddress
	goth.UseProviders(
		google.New(server.config.GoogleKey, server.config.GoogleSecret, BaseURL+"/auth/google/callback"),
	)

	router := gin.Default()
	m := map[string]string{
		"google": "Google",
	}
	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	_ = &ProviderIndex{Providers: keys, ProvidersMap: m}

	api := router.Group("/api/v1")

	api.GET("/health", server.healthCheck)

	auth := api.Group("/auth")
	auth.POST("/register", server.registerUser)
	auth.POST("/login", server.loginUser)
	auth.GET("/:provider/callback", server.providerCallback)
	auth.GET("/:provider", server.providerLogin)

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
