package handlers

import (
  // "ecom-auth/models"
  "ecom-auth/pkg/utils"
  "ecom-auth/repository"
  "errors"
  // "github.com/gin-gonic/gin"
  // "log"
  // "net/http"
  // "fmt"
)

// func (auth *AuthController) Login(c *gin.Context) {
// 	var userlogin models.Users
// 	err := c.BindJSON(&userlogin)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	var userinfor models.Users
// 	if err := auth.db.Where("username = ?", userlogin.Username).First(&userinfor).Error; err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "Invalid username"})
// 		return
// 	}

// 	PasswordIsValid, msg := utils.VerifyPassword(*userlogin.Password, *userinfor.Password)
// 	if !PasswordIsValid {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
// 		fmt.Println(msg)
// 		return
// 	}

// 	token, refreshToken, _ := utils.TokenGenerator(*userinfor.Email, *userinfor.Username, *userinfor.Fullname)

// 	if err := auth.db.Model(&userinfor).Where("user_id = ?", *userinfor.UserID).Updates(map[string]interface{}{"token": token, "refresh_token": refreshToken}).Error; err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "Error update token"})
// 		log.Println(err)
// 		return
// 	}
// 	c.JSON(http.StatusFound, userinfor.Token)
// }


func AuthenticateUser(username, password string) (string, error) {
	user, err := repository.GetUserByUsername(username)
	if err != nil {
		return "", errors.New("user not found")
	}

	if !utils.CheckPasswordHash(password, user.Password) {
		return "", errors.New("invalid password")
	}

	// Tạo JWT token
	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		return "", errors.New("failed to generate token")
	}

	return token, nil
}