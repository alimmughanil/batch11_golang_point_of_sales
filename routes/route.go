package routes

import (
	AuthController "batch11_golang_pos/controllers/auth"
	UserController "batch11_golang_pos/controllers/user"
	"os"

	echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitRoutes(e *echo.Echo) {
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.POST("/login", AuthController.Login)
	e.POST("/register", AuthController.Register)

	AuthGroup := e.Group("")
	AuthGroup.Use(echojwt.JWT([]byte(os.Getenv("JWT_SECRET"))))
	AuthGroup.GET("/users", UserController.GetAllUser)
	AuthGroup.GET("/users/:id", UserController.GetFirstUser)
}
