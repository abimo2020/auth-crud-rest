package controllers

import (
	"net/http"

	"github.com/Auth-CRUD-REST/models/payload"
	"github.com/Auth-CRUD-REST/usecase"
	"github.com/Auth-CRUD-REST/util"
	uuid "github.com/satori/go.uuid"

	"github.com/labstack/echo/v4"
)

func GetUsersController(c echo.Context) error {
	keyword := c.QueryParam("keyword")
	email := c.QueryParam("email")
	users, err := usecase.GetUsers(keyword, email)

	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.NewHTTPError(http.StatusBadRequest, err.Error()))
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all users",
		"data":    users,
	})
}

func GetUserByIdController(c echo.Context) error {
	id := c.Param("id")
	uuid, _ := uuid.FromString(id)
	user, err := usecase.GetUserById(uuid)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.NewHTTPError(http.StatusBadRequest, err.Error()))
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get user",
		"user":    user,
	})
}
func CreateUserController(c echo.Context) error {
	var req payload.UserRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.NewHTTPError(http.StatusBadRequest, err.Error()))
	}

	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.NewHTTPError(http.StatusBadRequest, err.Error()))
	}

	if err := usecase.CreateUser(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.NewHTTPError(http.StatusBadRequest, err.Error()))
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success to create user",
	})
}
func UpdateUserController(c echo.Context) error {
	var req payload.UserRequest
	id := c.Param("id")
	uuid, _ := uuid.FromString(id)

	c.Bind(&req)

	authId := util.GetUserLoginId(c)
	authRole := util.GetUserRole(c)
	if authId != id && authRole != "admin" {
		return c.JSON(http.StatusBadRequest, echo.NewHTTPError(http.StatusBadRequest, "Must be admin or owner account"))
	}

	if err := usecase.UpdateUser(&req, uuid); err != nil {
		return c.JSON(http.StatusBadRequest, echo.NewHTTPError(http.StatusBadRequest, err.Error()))
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update user",
	})
}

func DeleteUserController(c echo.Context) error {
	id := c.Param("id")
	uuid, _ := uuid.FromString(id)
	if err := usecase.DeleteUser(uuid); err != nil {
		return c.JSON(http.StatusBadRequest, echo.NewHTTPError(http.StatusBadRequest, err.Error()))
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete user",
	})
}
