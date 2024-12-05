package controller

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/mhdianrush/ecommerce-project/config"
	"github.com/mhdianrush/ecommerce-project/entities"
	"github.com/mhdianrush/ecommerce-project/request"
	"github.com/mhdianrush/ecommerce-project/response"
)

func CreateBrand(c echo.Context) error {
	var req request.CreateBrandRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	brand := entities.Brands{
		NamaBrand: req.NamaBrand,
	}

	if err := config.DB.Create(&brand).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create brand"})
	}

	resp := response.CreateBrandResponse{
		ID: brand.IdBrand,
	}

	return c.JSON(http.StatusOK, resp)
}

func DeleteBrand(c echo.Context) error {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid brand ID"})
	}

	var count int64
	if err := config.DB.Model(&entities.Products{}).Where("id_brand = ?", idInt).Count(&count).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to check brand usage"})
	}

	if count > 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Cannot delete brand, it is in use by products"})
	}

	if err := config.DB.Delete(&entities.Brands{}, idInt).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to delete brand"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Brand deleted successfully"})
}
