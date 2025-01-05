package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok",
		"message": "System is healthy",
	})

}


func (s *Server) registerUser(c *gin.Context) {

}
