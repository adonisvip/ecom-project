package utils

import (
	"errors"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type SignedDetails struct {
	Email    string
	Username string
	Fullname string
	jwt.StandardClaims
}

var SECRET_KEY = os.Getenv("SECRET_KEY")

// TokenGenerator generates access token (JWT) with 15 minutes expiration
func TokenGenerator(email string, username string, fullname string) (signedtoken string, err error) {
	claims := &SignedDetails{
		Email:    email,
		Username: username,
		Fullname: fullname,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Minute * 15).Unix(),
		},
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))
	if err != nil {
		return "", err
	}
	return token, nil
}

func ValidateToken(signedtoken string) (claims *SignedDetails, msg string) {
	token, err := jwt.ParseWithClaims(signedtoken, &SignedDetails{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SECRET_KEY), nil
	})

	if err != nil {
		msg = err.Error()
		return
	}
	claims, ok := token.Claims.(*SignedDetails)
	if !ok {
		msg = "The Token is invalid"
		return
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		msg = "token is expired"
		return
	}
	return claims, msg
}

func ParseJWT(tokenString string) (jwt.MapClaims, error) {
	if SECRET_KEY == "" {
		return nil, errors.New("SECRET_KEY is not set")
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(SECRET_KEY), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

func GenerateJWT(userID int) (string, error) {
	if SECRET_KEY == "" {
		return "", errors.New("SECRET_KEY is not set")
	}

	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(SECRET_KEY))
}

func GenerateTokenPair(userID int) (token string, refreshToken string, err error) {
	if SECRET_KEY == "" {
		return "", "", errors.New("SECRET_KEY is not set")
	}

	// Generate access token (24 hours)
	tokenClaims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}
	tokenObj := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims)
	token, err = tokenObj.SignedString([]byte(SECRET_KEY))
	if err != nil {
		return "", "", err
	}

	// Generate refresh token (7 days)
	refreshClaims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24 * 7).Unix(),
	}
	refreshTokenObj := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshToken, err = refreshTokenObj.SignedString([]byte(SECRET_KEY))
	if err != nil {
		return "", "", err
	}

	return token, refreshToken, nil
}
