package redis

import (
  "time"
  "context"
  "log"
  "github.com/go-redis/redis/v8"
  "ecom-auth/config"
)

func InitConnection(cfg *config.Redis) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:         cfg.ConnectionURL,
		PoolSize:     cfg.PoolSize,
		PoolTimeout:  time.Duration(cfg.PoolTimeout) * time.Second,
		Password:     cfg.Password,
		DB:           cfg.DB,  
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
	return client
}