package handlers

import (
	// "ecom-auth/models"
	// "ecom-auth/pkg/utils"
	"ecom-auth/repository"
	"errors"
	// "net/http"
	// "time"
	// "github.com/gin-gonic/gin"
)

// func (auth *AuthController) SignUp(c *gin.Context) {
// 	var userinfo models.Users
// 	err := c.BindJSON(&userinfo)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	validationErr := Validate.Struct(userinfo)
// 	if validationErr != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": validationErr})
// 		return
// 	}

// 	countEmailExist := auth.db.Where("email = ?", userinfo.Email).First(&models.Users{}).RowsAffected
// 	if countEmailExist > 0 {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Email already exist"})
// 		return
// 	}

// 	countPhoneExist := auth.db.Where("phone = ?", userinfo.Phone).First(&models.Users{}).RowsAffected
// 	if countPhoneExist > 0 {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Phone already exist"})
// 		return
// 	}

// 	countUserExist := auth.db.Where("username = ?", userinfo.Username).First(&models.Users{}).RowsAffected
// 	if countUserExist > 0 {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Username already exist"})
// 		return
// 	}

// 	password := utils.HashPass(*userinfo.Password)
// 	userinfo.Password = &password
// 	userinfo.Created_At, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
// 	userinfo.Updated_At, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
// 	token, refreshtoken, _ := utils.TokenGenerator(*userinfo.Email, *userinfo.Username, *userinfo.Fullname)
// 	userinfo.Token = &token
// 	userinfo.Refresh_Token = &refreshtoken

// 	if err := auth.db.Create(&userinfo).Error; err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	c.JSON(http.StatusCreated, gin.H{"message": "Signup successful"})
// }

// RegisterUser handles user registration via gRPC
func RegisterUser(username, password, email string) error {
	// Check if username already exists
	exists, err := repository.UserExists(username)
	if err != nil {
		return errors.New("failed to check username availability")
	}
	if exists {
		return errors.New("username already exists")
	}

	// Check if email already exists
	emailExists, err := repository.EmailExists(email)
	if err != nil {
		return errors.New("failed to check email availability")
	}
	if emailExists {
		return errors.New("email already exists")
	}

	// Create the user
	err = repository.CreateUser(username, password, email)
	if err != nil {
		return errors.New("failed to create user")
	}

	return nil
}
