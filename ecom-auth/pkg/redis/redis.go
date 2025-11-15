package redis

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
	"ecom-auth/config"
)

var Client *redis.Client

func InitConnection(cfg *config.Redis) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:        cfg.ConnectionURL,
		PoolSize:    cfg.PoolSize,
		PoolTimeout: time.Duration(cfg.PoolTimeout) * time.Second,
		Password:    cfg.Password,
		DB:          cfg.DB,
	})
	// Test connection with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := client.Ping(ctx).Result()

	if err != nil {
		log.Printf("Failed to connect to Redis at %s: %v", cfg.ConnectionURL, err)
		return nil
	}

	log.Printf("Successfully connected to Redis at %s", cfg.ConnectionURL)
	Client = client
	return client
}

// RefreshTokenData stores user info for refresh token
type RefreshTokenData struct {
	UserID   int    `json:"user_id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Fullname string `json:"fullname"`
}

// SaveRefreshToken saves refresh token to Redis with 7 days expiration
func SaveRefreshToken(refreshToken string, data RefreshTokenData) error {
	ctx := context.Background()
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	key := "refresh_token:" + refreshToken
	return Client.Set(ctx, key, jsonData, 7*24*time.Hour).Err()
}

// GetRefreshToken gets refresh token data from Redis
func GetRefreshToken(refreshToken string) (*RefreshTokenData, error) {
	ctx := context.Background()
	key := "refresh_token:" + refreshToken

	val, err := Client.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	var data RefreshTokenData
	err = json.Unmarshal([]byte(val), &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

// DeleteRefreshToken deletes refresh token from Redis
func DeleteRefreshToken(refreshToken string) error {
	ctx := context.Background()
	key := "refresh_token:" + refreshToken
	return Client.Del(ctx, key).Err()
}