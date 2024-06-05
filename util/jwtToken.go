package util

import (
	"context"
	"os"
	"time"

	"github.com/Auth-CRUD-REST/config"
	"github.com/go-redis/redis/v8"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"
)

func GenerateAndStoreToken(email, role string, id uuid.UUID) (string, error) {
	token, err := config.Redis.Get(context.Background(), email).Result()
	if err == redis.Nil {
		claims := jwt.MapClaims{}
		claims["email"] = email
		claims["role"] = role
		claims["user_id"] = id
		claims["authorized"] = true
		claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
		jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		token, err = jwtToken.SignedString([]byte(os.Getenv("JWT_SECRET")))
		if err != nil {
			return "", err
		}
		err = config.Redis.Set(context.Background(), email, token, 24*time.Hour).Err()
		if err != nil {
			return "", err
		}
	} else if err != nil {
		return "", err
	}

	return token, nil
}

func GetUserLoginId(c echo.Context) string {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	id := claims["user_id"].(string)
	return id
}

func GetUserRole(c echo.Context) string {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	role := claims["role"].(string)
	return role
}
