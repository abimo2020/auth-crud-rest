package database

import (
	"github.com/Auth-CRUD-REST/config"
	"github.com/Auth-CRUD-REST/models"
	"github.com/Auth-CRUD-REST/models/payload"
	uuid "github.com/satori/go.uuid"
)

func CreateUser(user *models.User) error {
	if err := config.DB.Save(&user).Error; err != nil {
		return err
	}
	return nil
}

func GetUsers(req *payload.UserParam) ([]models.User, error) {
	var user []models.User
	db := config.DB

	if req.Keyword != "" {
		db = db.Where("name like ?", "%"+req.Keyword+"%")
	}
	if req.Email != "" {
		db = db.Where("email like ?", "%"+req.Email+"%")
	}
	if err := db.Not("role = ?", "admin").Find(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func GetUserById(id uuid.UUID) (resp models.User, err error) {
	if err := config.DB.Where("id = ?", id).Not("role = ?", "admin").First(&resp).Error; err != nil {
		return resp, err
	}
	return
}

func GetUserByEmail(email string) (user models.User, err error) {
	if err := config.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return models.User{}, err
	}
	return
}

func UpdateUser(user *models.User, id uuid.UUID) error {
	if err := config.DB.Where("id = ?", id).Updates(&user).Error; err != nil {
		return err
	}
	return nil
}
func DeleteUser(id uuid.UUID) error {
	user := models.User{}
	if err := config.DB.Model(&user).Where("id = ?", id).Delete(&user).Error; err != nil {
		return err
	}
	return nil
}
