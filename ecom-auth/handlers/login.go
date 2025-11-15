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

func AuthenticateUser(username, password string) (string, string, error) {
	user, err := repository.GetUserByUsername(username)
	if err != nil {
		return "", "", errors.New("user not found")
	}

	if !utils.CheckPasswordHash(password, user.Password) {
		return "", "", errors.New("invalid password")
	}

	// Lấy fullname, nếu không có thì dùng empty string
	fullname := ""
	if user.Fullname.Valid {
		fullname = user.Fullname.String
	}

	// Tạo access token (JWT) với thời hạn 15 phút
	token, err := utils.TokenGenerator(user.Email, user.Username, fullname)
	if err != nil {
		return "", "", errors.New("failed to generate token")
	}

	// Tạo refresh token (UUID) với thời hạn 14 ngày (lưu trong DB)
	refreshToken := utils.GenerateRefreshToken()

	// Lưu cả access token và refresh token vào database
	err = repository.UpdateUserTokens(user.ID, token, refreshToken)
	if err != nil {
		return "", "", errors.New("failed to save token to database")
	}

	return token, refreshToken, nil
}
