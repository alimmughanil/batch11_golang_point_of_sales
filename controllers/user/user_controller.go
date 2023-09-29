package UserController

import (
	"batch11_golang_pos/configs"
	UserModel "batch11_golang_pos/models/user"
	"batch11_golang_pos/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetAllUser(c echo.Context) error {
	var users []UserModel.User
	result := configs.DB.Find(&users)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, utils.Response{
			Status:  false,
			Message: "Failed get data from database",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, utils.Response{
		Status:  true,
		Message: "Success",
		Data:    users,
	})
}
func GetFirstUser(c echo.Context) error {
	var user UserModel.User
	c.Bind(&user)
	result := configs.DB.Find(&user)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, utils.Response{
			Status:  false,
			Message: "Failed get data from database",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, utils.Response{
		Status:  true,
		Message: "Success",
		Data:    user,
	})
}
