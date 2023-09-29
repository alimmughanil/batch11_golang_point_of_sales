package AuthController

import (
	"batch11_golang_pos/configs"
	"batch11_golang_pos/middleware"
	UserModel "batch11_golang_pos/models/user"
	"batch11_golang_pos/utils"
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func Register(c echo.Context) error {
	var user UserModel.User
	c.Bind(&user)

	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 12)
	user.Password = string(hashPassword)

	result := configs.DB.Create(&user)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, utils.Response{
			Status:  false,
			Message: "Failed add data to database",
			Data:    nil,
		})
	}

	var authResponse = UserModel.Response{
		Id:    user.Id,
		Name:  user.Name,
		Email: user.Email,
		Token: middleware.GenerateTokenJWT(user.Id, user.Name),
	}

	return c.JSON(http.StatusOK, utils.Response{
		Status:  true,
		Message: "Success create new account",
		Data:    authResponse,
	})
}

func Login(c echo.Context) error {
	var user UserModel.User
	c.Bind(&user)
	password := user.Password
	result := configs.DB.Find(&user)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, utils.Response{
			Status:  false,
			Message: "Failed get to database",
			Data:    nil,
		})
	}

	validatePassword := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if validatePassword != nil {
		return c.JSON(http.StatusBadRequest, utils.Response{
			Status:  false,
			Message: "Password is invalid",
			Data:    nil,
		})
	}

	var authResponse = UserModel.Response{
		Id:    user.Id,
		Name:  user.Name,
		Email: user.Email,
		Token: middleware.GenerateTokenJWT(user.Id, user.Name),
	}

	return c.JSON(http.StatusOK, utils.Response{
		Status:  true,
		Message: "Success get authentication token",
		Data:    authResponse,
	})
}
