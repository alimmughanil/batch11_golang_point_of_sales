package CategoryController

import (
	"batch11_golang_pos/configs"
	CategoryModel "batch11_golang_pos/models/category"
	"batch11_golang_pos/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetAllCategory(c echo.Context) error {
	var catagory []CategoryModel.Category
	result := configs.DB.Find(&catagory)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, utils.Response{
			Status:  false,
			Message: "Failed get data from database",
			Data:    nil,
		})
	}
	if result.RowsAffected == 0 {
		return c.JSON(http.StatusNotFound, utils.Response{
			Status:  false,
			Message: "Data Not Found",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, utils.Response{
		Status:  true,
		Message: "Success",
		Data:    catagory,
	})
}

func GetFirstCategory(c echo.Context) error {
	var category CategoryModel.Category
	id := c.Param("id")
	result := configs.DB.Model(&category).Where("id = ?", id).Find(&category)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, utils.Response{
			Status:  false,
			Message: "Failed get data from database",
			Data:    nil,
		})
	}
	if result.RowsAffected == 0 {
		return c.JSON(http.StatusNotFound, utils.Response{
			Status:  false,
			Message: "Data Not Found",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, utils.Response{
		Status:  true,
		Message: "Success",
		Data:    category,
	})
}

func CreateCategory(c echo.Context) error {
	var category CategoryModel.Category
	c.Bind(&category)
	result := configs.DB.Create(&category)

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
		Data:    category,
	})
}

func UpdateCategory(c echo.Context) error {
	var category CategoryModel.Category
	c.Bind(&category)
	id := c.Param("id")
	title := category.Title
	result := configs.DB.Model(&category).Where("id = ?", id).Update("title", title)
	category.Id, _ = strconv.Atoi(id)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, utils.Response{
			Status:  false,
			Message: "Failed get data from database",
			Data:    nil,
		})
	}
	if result.RowsAffected == 0 {
		return c.JSON(http.StatusNotFound, utils.Response{
			Status:  false,
			Message: "Data Not Found",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, utils.Response{
		Status:  true,
		Message: "Success",
		Data:    category,
	})
}

func DeleteCategory(c echo.Context) error {
	var category CategoryModel.Category
	id := c.Param("id")
	result := configs.DB.Model(&category).Where("id = ?", id).Delete(&category)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, utils.Response{
			Status:  false,
			Message: "Failed get data from database",
			Data:    nil,
		})
	}
	if result.RowsAffected == 0 {
		return c.JSON(http.StatusNotFound, utils.Response{
			Status:  false,
			Message: "Data Not Found",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, utils.Response{
		Status:  true,
		Message: "Success",
		Data:    nil,
	})
}
