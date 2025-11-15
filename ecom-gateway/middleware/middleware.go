package middleware

import (
	"context"
	"net/http"
	"strings"

	"ecom-gateway/grpc"
	pbAuth "ecom-gateway/proto/auth"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware kiểm tra JWT token
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
			c.Abort()
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")
		if token == authHeader {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
			c.Abort()
			return
		}

		res, err := grpc.AuthClient.ValidateToken(context.Background(), &pbAuth.ValidateTokenRequest{Token: token})
		if err != nil || !res.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		c.Next()
	}

}
