package usecase

import (
	"net/http"

	"github.com/Auth-CRUD-REST/models"
	"github.com/Auth-CRUD-REST/models/payload"
	"github.com/Auth-CRUD-REST/repository/database"
	"github.com/Auth-CRUD-REST/util"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func Register(req *payload.UserRequest) (string, error) {
	if req.Password != req.RetypePassword {
		return "", echo.NewHTTPError(http.StatusBadRequest, "Password not match")
	}
	password, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	user := models.User{
		Email:       req.Email,
		Password:    string(password),
		Name:        req.Name,
		PhoneNumber: req.PhoneNumber,
		Address:     req.PhoneNumber,
	}
	if err := database.CreateUser(&user); err != nil {
		return "", err
	}
	token, err := util.GenerateAndStoreToken(user.Email, user.Role, user.ID)
	if err != nil {
		return "", err
	}
	return token, nil
}

func Login(req *payload.LoginRequest) (string, error) {
	user, err := database.GetUserByEmail(req.Email)
	if err != nil {
		return "", err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return "", err
	}
	token, err := util.GenerateAndStoreToken(user.Email, user.Role, user.ID)
	if err != nil {
		return "", err
	}
	return token, nil
}
