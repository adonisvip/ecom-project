package handlers

import (
  "github.com/go-playground/validator/v10"
  "gorm.io/gorm"
  "ecom-auth/pkg/utils"
)

type AuthController struct {
	db *gorm.DB
}

func NewAuthController(db *gorm.DB) AuthController {
	return AuthController{db}
}

var Validate = validator.New()

func ValidateJWT(token string) bool {
	_, err := utils.ParseJWT(token)
	return err == nil
}