package routes

import (
	"os"

	"github.com/Auth-CRUD-REST/controllers"
	m "github.com/Auth-CRUD-REST/middlewares"
	jwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	e := echo.New()

	m.LoggerMiddleware(e)
	e.Pre(middleware.RemoveTrailingSlash())

	e.POST("/register", controllers.RegisterController)
	e.POST("/login", controllers.LoginController)

	users := e.Group("/users")
	users.GET("", controllers.GetUsersController)
	users.GET("/:id", controllers.GetUserByIdController)
	users.POST("", controllers.CreateUserController, jwt.JWT([]byte(os.Getenv("JWT_SECRET"))), m.IsAdmin)
	users.PUT("/:id", controllers.UpdateUserController, jwt.JWT([]byte(os.Getenv("JWT_SECRET"))))
	users.DELETE("/:id", controllers.DeleteUserController, jwt.JWT([]byte(os.Getenv("JWT_SECRET"))), m.IsAdmin)

	return e
}
