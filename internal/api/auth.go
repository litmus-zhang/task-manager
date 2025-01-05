package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/litmus-zhang/task-manager/internal/db"
	"github.com/litmus-zhang/task-manager/util"
)

func (s *Server) healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok",
		"message": "System is healthy",
	})

}

type registerUserRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	FullName string `json:"full_name" binding:"required"`
	UserName string `json:"username" binding:"required"`
}

func (s *Server) registerUser(c *gin.Context) {
	var req registerUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		errResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	hash, err := util.HashPassword(req.Password)
	if err != nil {
		errResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	args := db.RegisterUserParams{
		Username:     req.UserName,
		Email:        req.Email,
		PasswordHash: hash,
		FullName:     req.FullName,
	}
	user, err := s.store.RegisterUser(c, args)
	if err != nil {
		errResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	user.PasswordHash = ""
	c.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
		"user":    user,
	})

}
