package handlers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"ecom-gateway/grpc"
	pbAuth "ecom-gateway/proto/auth"
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
