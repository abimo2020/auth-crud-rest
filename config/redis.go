package config

import (
	"context"
	"fmt"
	"os"

	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
)

var Redis *redis.Client

type RedisConfig struct {
	Host     string
	Password string
	DB       int
}

func RedisInit() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file")
	}
	config := RedisConfig{
		Host:     os.Getenv("REDIS_HOST"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	}
	Redis = redis.NewClient(&redis.Options{
		Addr:     config.Host,
		Password: config.Password,
		DB:       config.DB,
	})
	_, err := Redis.Ping(context.Background()).Result()
	if err != nil {
		fmt.Println("Failed to connect to Redis:", err)
	} else {
		fmt.Println("Connected to Redis!")
	}
}
