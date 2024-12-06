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
		return c.JSON(http.StatusBadRequest, response.TemplateResponse{
			Message: "invalid request",
			Data:    nil,
		})
	}

	brand := entities.Brands{
		NamaBrand: req.NamaBrand,
	}

	tx := config.DB.Begin()
	if err := tx.Create(&brand).Error; err != nil {
		tx.Rollback()
		return c.JSON(http.StatusInternalServerError, response.TemplateResponse{
			Message: "failed to create brand",
			Data:    nil,
		})
	}
	tx.Commit()

	resp := response.CreateBrandResponse{
		ID: brand.IdBrand,
	}

	return c.JSON(http.StatusOK, response.TemplateResponse{
		Message: "success insert new brand",
		Data:    resp,
	})
}

func DeleteBrand(c echo.Context) error {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.TemplateResponse{
			Message: "invalid brand ID",
			Data:    nil,
		})
	}

	var count int64
	if err := config.DB.Model(&entities.Products{}).Where("id_brand = ?", idInt).Count(&count).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, response.TemplateResponse{
			Message: "failed to check brand usage",
			Data:    nil,
		})
	}

	if count > 0 {
		return c.JSON(http.StatusBadRequest, response.TemplateResponse{
			Message: "can't delete brand, it is in use by products",
			Data:    nil,
		})
	}

	tx := config.DB.Begin()
	if err := tx.Delete(&entities.Brands{}, idInt).Error; err != nil {
		tx.Rollback()
		return c.JSON(http.StatusInternalServerError, response.TemplateResponse{
			Message: "failed to delete brand",
			Data:    nil,
		})
	}
	tx.Commit()

	return c.JSON(http.StatusOK, response.TemplateResponse{
		Message: "brand deleted successfully",
		Data:    nil,
	})
}
