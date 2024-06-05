package usecase

import (
	"encoding/json"
	"net/http"

	"github.com/Auth-CRUD-REST/models"
	"github.com/Auth-CRUD-REST/repository/database"
	"github.com/Auth-CRUD-REST/services"
	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"

	"github.com/Auth-CRUD-REST/models/payload"
)

func GetUsers(keyword, email string) (response []payload.UserResponse, err error) {
	req := payload.UserParam{
		Keyword: keyword,
		Email:   email,
	}
	users, err := database.GetUsers(&req)

	for _, user := range users {
		response = append(response, payload.UserResponse{
			ID:          user.ID,
			Name:        user.Name,
			Email:       user.Email,
			PhoneNumber: user.PhoneNumber,
			Address:     user.Address,
			Role:        user.Role,
		})
	}
	return
}

func GetUserById(id uuid.UUID) (resp payload.UserResponse, err error) {
	user, err := database.GetUserById(id)
	if err != nil {
		return payload.UserResponse{}, err
	}
	resp = payload.UserResponse{
		ID:          user.ID,
		Name:        user.Name,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		Address:     user.Address,
		Role:        user.Role,
	}
	return
}

func CreateUser(req *payload.UserRequest) error {
	password, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user := models.User{
		Email:       req.Email,
		Password:    string(password),
		Name:        req.Name,
		PhoneNumber: req.PhoneNumber,
		Address:     req.PhoneNumber,
		Role:        req.Role,
	}
	if err := database.CreateUser(&user); err != nil {
		return err
	}
	data, err := json.Marshal(req)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	message := services.Message{
		Action: "CreateUser",
		Data:   string(data),
	}

	if err := services.PublishMessage(message); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to publish message")
	}
	return nil
}

func UpdateUser(req *payload.UserRequest, id uuid.UUID) (err error) {
	if _, err := database.GetUserById(id); err != nil {
		return err
	}
	user := models.User{
		Name:        req.Name,
		PhoneNumber: req.PhoneNumber,
		Address:     req.Address,
		Password:    req.Password,
	}
	if err := database.UpdateUser(&user, id); err != nil {
		return err
	}
	data, err := json.Marshal(req)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	message := services.Message{
		Action: "UpdateUser",
		Data:   string(data),
	}

	if err := services.PublishMessage(message); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to publish message")
	}
	return
}

func DeleteUser(id uuid.UUID) (err error) {
	if _, err := database.GetUserById(id); err != nil {
		return err
	}
	err = database.DeleteUser(id)
	if err != nil {
		return
	}
	return
}
