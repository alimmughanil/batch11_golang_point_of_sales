package ProductController

import (
	"batch11_golang_pos/configs"
	ProductModel "batch11_golang_pos/models/product"
	"batch11_golang_pos/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetAllProduct(c echo.Context) error {
	var catagory []ProductModel.Product
	result := configs.DB.Preload("Category").Preload("Category").Find(&catagory)

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

func GetFirstProduct(c echo.Context) error {
	var product ProductModel.Product
	id := c.Param("id")
	result := configs.DB.Preload("Category").Model(&product).Where("id = ?", id).Find(&product)

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
		Data:    product,
	})
}

func CreateProduct(c echo.Context) error {
	var product ProductModel.Product
	c.Bind(&product)
	result := configs.DB.Create(&product)

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
		Data:    product,
	})
}

func UpdateProduct(c echo.Context) error {
	var product ProductModel.Product
	c.Bind(&product)
	id := c.Param("id")
	result := configs.DB.Model(&product).Where("id = ?", id).Updates(&product)
	product.Id, _ = strconv.Atoi(id)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, utils.Response{
			Status:  false,
			Message: "Failed get data from database",
			Data:    nil,
		})
	}
	if result.RowsAffected == 0 {
		return c.JSON(http.StatusAccepted, utils.Response{
			Status:  false,
			Message: "No data updated",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, utils.Response{
		Status:  true,
		Message: "Success",
		Data:    product,
	})
}

func DeleteProduct(c echo.Context) error {
	var product ProductModel.Product
	id := c.Param("id")
	result := configs.DB.Model(&product).Where("id = ?", id).Delete(&product)

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
