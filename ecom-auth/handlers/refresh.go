package handlers

import (
	"ecom-auth/pkg/utils"
	"ecom-auth/repository"
	"errors"
)

// RefreshToken generates a new access token and refresh token using old refresh token
func RefreshToken(refreshToken string) (string, string, error) {
	// Verify refresh token trong DB
	user, err := repository.GetUserByRefreshToken(refreshToken)
	if err != nil {
		return "", "", errors.New("refresh token not found or invalid")
	}

	// Lấy fullname, nếu không có thì dùng empty string
	fullname := ""
	if user.Fullname.Valid {
		fullname = user.Fullname.String
	}

	// Tạo access token mới (JWT) với thời hạn 15 phút
	newAccessToken, err := utils.TokenGenerator(user.Email, user.Username, fullname)
	if err != nil {
		return "", "", errors.New("failed to generate access token")
	}

	// Tạo refresh token mới (UUID) với thời hạn 14 ngày
	newRefreshToken := utils.GenerateRefreshToken()

	// Xóa RT cũ và lưu RT mới + access token mới vào database
	err = repository.UpdateUserTokens(user.ID, newAccessToken, newRefreshToken)
	if err != nil {
		return "", "", errors.New("failed to update tokens in database")
	}

	return newAccessToken, newRefreshToken, nil
}

