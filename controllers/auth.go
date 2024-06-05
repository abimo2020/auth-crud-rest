package controllers

import (
	"net/http"

	"github.com/Auth-CRUD-REST/models/payload"
	"github.com/Auth-CRUD-REST/usecase"
	"github.com/labstack/echo/v4"
)

func RegisterController(c echo.Context) error {
	var req payload.UserRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.NewHTTPError(http.StatusBadRequest, err.Error()))
	}

	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.NewHTTPError(http.StatusBadRequest, err.Error()))
	}

	token, err := usecase.Register(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.NewHTTPError(http.StatusBadRequest, err.Error()))
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success to register",
		"token":   token,
	})
}

func LoginController(c echo.Context) error {
	var req payload.LoginRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.NewHTTPError(http.StatusBadRequest, err.Error()))
	}

	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.NewHTTPError(http.StatusBadRequest, err.Error()))
	}

	token, err := usecase.Login(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.NewHTTPError(http.StatusBadRequest, err.Error()))
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success login user",
		"token":   token,
	})
}
