package models

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type User struct {
	Base
	Email       string `json:"email" form:"email" gorm:"unique"`
	Password    string `json:"password" form:"password"`
	Role        string `json:"role" form:"role" gorm:"type:enum('user','admin');default:'user'"`
	Name        string `json:"name" form:"name"`
	PhoneNumber string `json:"phone_number" form:"phone_number"`
	Address     string `json:"address" form:"address"`
}

func (user *User) BeforeDelete(tx *gorm.DB) error {
	if user.Role == "admin" {
		return echo.NewHTTPError(http.StatusBadRequest, "Cant delete admin user")
	}
	return nil
}
