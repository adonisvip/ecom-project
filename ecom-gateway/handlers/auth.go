package handlers

import (
	"context"
	"net/http"

	"ecom-gateway/grpc"
	pbAuth "ecom-gateway/proto/auth"

	"github.com/gin-gonic/gin"
)

func LoginHandler(c *gin.Context) {
	var req pbAuth.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	res, err := grpc.AuthClient.Login(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Login failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": res.Token})
}

func SignupHandler(c *gin.Context) {
	var req pbAuth.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	res, err := grpc.AuthClient.Register(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Registration failed"})
		return
	}

	if res.Message != "Registration successful" {
		c.JSON(http.StatusBadRequest, gin.H{"error": res.Message})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": res.Message})
}
