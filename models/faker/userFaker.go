package faker

import (
	"github.com/Auth-CRUD-REST/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func UserFaker(db *gorm.DB) *models.User {
	password, _ := bcrypt.GenerateFromPassword([]byte("admin"), bcrypt.DefaultCost)
	return &models.User{
		Email:    "admin",
		Password: string(password),
		Role:     "admin",
	}
}
