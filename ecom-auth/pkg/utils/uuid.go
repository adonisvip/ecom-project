package utils

import (
	"github.com/google/uuid"
)

// GenerateRefreshToken generates a UUID string for refresh token
func GenerateRefreshToken() string {
	return uuid.New().String()
}

