package redis

import (
  "time"
  "context"
  "log"
  "github.com/go-redis/redis/v8"
  "ecom-gateway/config"
)

func InitConnection(cfg *config.Config) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:         cfg.Redis.ConnectionURL,
		PoolSize:     cfg.Redis.PoolSize,
		PoolTimeout:  time.Duration(cfg.Redis.PoolTimeout) * time.Second,
		Password:     cfg.Redis.Password,
		DB:           cfg.Redis.DB,  
	})
  	// Test connection with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := client.Ping(ctx).Result()

	if err != nil {
		log.Printf("Failed to connect to Redis at %s: %v", cfg.Redis.ConnectionURL, err)
		return nil
	}

	log.Printf("Successfully connected to Redis at %s", cfg.Redis.ConnectionURL)
	return client
}