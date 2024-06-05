package middlewares

import (
	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func IsAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		isAdmin := claims["role"].(string)
		if isAdmin != "admin" {
			return echo.ErrUnauthorized
		}
		return next(c)
	}
}

func IsUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		isUser := claims["role"].(string)
		if isUser == "admin" {
			return echo.ErrUnauthorized
		}
		return next(c)
	}
}
