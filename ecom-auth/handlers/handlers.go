package handlers

import (
	"ecom-auth/pkg/utils"
	"ecom-auth/repository"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type AuthController struct {
	db *gorm.DB
}

func NewAuthController(db *gorm.DB) AuthController {
	return AuthController{db}
}

var Validate = validator.New()

func ValidateJWT(token string) bool {
	// First validate JWT signature and expiration using ValidateToken (compatible with TokenGenerator)
	_, msg := utils.ValidateToken(token)
	if msg != "" {
		return false
	}

	// Then check if token exists in database
	_, err := repository.GetUserByToken(token)
	return err == nil
}