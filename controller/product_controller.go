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

func CreateProduct(c echo.Context) error {
	var req request.CreateProductRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, response.TemplateResponse{
			Message: "invalid request payload",
			Data:    nil,
		})
	}

	product := entities.Products{
		NamaProduct: req.NamaProduct,
		Harga:       req.Harga,
		Quantity:    req.Quantity,
		IdBrand:     req.IdBrand,
	}

	if err := config.DB.Create(&product).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, response.TemplateResponse{
			Message: "failed to insert product",
			Data:    nil,
		})
	}

	resp := response.CreateProductResponse{ID: product.ID}

	return c.JSON(http.StatusCreated, response.TemplateResponse{
		Message: "success insert new product",
		Data:    resp,
	})
}

func GetProductById(c echo.Context) error {
	id := c.Param("id")

	productID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.TemplateResponse{
			Message: "invalid product id",
			Data:    nil,
		})
	}

	var product entities.Products
	if err := config.DB.First(&product, productID).Error; err != nil {
		return c.JSON(http.StatusNotFound, response.TemplateResponse{
			Message: "product not found",
			Data:    nil,
		})
	}

	resp := response.GetProductByIdResponse{
		ID:          product.ID,
		NamaProduct: product.NamaProduct,
		Harga:       product.Harga,
		Quantity:    product.Quantity,
		IdBrand:     product.IdBrand,
	}

	return c.JSON(http.StatusOK, response.TemplateResponse{
		Message: "success get product detail",
		Data:    resp,
	})
}

func UpdateProduct(c echo.Context) error {
	id := c.Param("id")

	productID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.TemplateResponse{
			Message: "invalid product id",
			Data:    nil,
		})
	}

	var req request.UpdateProductRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, response.TemplateResponse{
			Message: "invalid request payload",
			Data:    nil,
		})
	}

	var product entities.Products
	if err := config.DB.First(&product, productID).Error; err != nil {
		return c.JSON(http.StatusNotFound, response.TemplateResponse{
			Message: "product not found",
			Data:    nil,
		})
	}

	product.NamaProduct = req.NamaProduct
	product.Harga = req.Harga
	product.Quantity = req.Quantity
	product.IdBrand = req.IdBrand

	if err := config.DB.Save(&product).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, response.TemplateResponse{
			Message: "failed to update product",
			Data:    nil,
		})
	}

	resp := response.UpdateProductResponse{
		NamaProduct: product.NamaProduct,
		Harga:       product.Harga,
		Quantity:    product.Quantity,
		IdBrand:     product.IdBrand,
	}

	return c.JSON(http.StatusOK, response.TemplateResponse{
		Message: "success update product data",
		Data:    resp,
	})
}

func DeleteProduct(c echo.Context) error {
	id := c.Param("id")

	productID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.TemplateResponse{
			Message: "invalid product id",
			Data:    nil,
		})
	}

	var product entities.Products
	if err := config.DB.First(&product, productID).Error; err != nil {
		return c.JSON(http.StatusNotFound, response.TemplateResponse{
			Message: "product not found",
			Data:    nil,
		})
	}

	if err := config.DB.Delete(&product).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, response.TemplateResponse{
			Message: "failed to delete product",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, response.TemplateResponse{
		Message: "product deleted successfully",
		Data:    nil,
	})
}

func GetAllProducts(c echo.Context) error {
	page := c.QueryParam("page")
	size := c.QueryParam("size")

	if page == "" {
		page = "1"
	}
	if size == "" {
		size = "10"
	}

	pageInt, _ := strconv.Atoi(page)
	sizeInt, _ := strconv.Atoi(size)

	offset := (pageInt - 1) * sizeInt

	var totalRecords int64
	if err := config.DB.Model(&entities.Products{}).Count(&totalRecords).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, response.TemplateResponse{
			Message: "failed to count products",
			Data:    nil,
		})
	}

	var products []entities.Products
	if err := config.DB.Offset(offset).Limit(sizeInt).Find(&products).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, response.TemplateResponse{
			Message: "failed to fetch products",
			Data:    nil,
		})
	}

	totalPages := int(totalRecords) / sizeInt
	if totalRecords%int64(sizeInt) != 0 {
		totalPages++
	}

	var productList []response.GetAllProductsDataResponse
	for _, product := range products {
		productList = append(productList, response.GetAllProductsDataResponse{
			ID:          product.ID,
			NamaProduct: product.NamaProduct,
			Harga:       product.Harga,
			Quantity:    product.Quantity,
			IdBrand:     product.IdBrand,
		})
	}

	pageInfo := response.PageInformation{
		CurrentPage: uint(pageInt),
		PageSize:    uint(sizeInt),
		Records:     uint(totalRecords),
		TotalPage:   uint(totalPages),
	}

	resp := response.GetAllProductsResponse{
		PageInformation: pageInfo,
		ListData:        productList,
	}

	return c.JSON(http.StatusOK, response.TemplateResponse{
		Message: "success get products",
		Data:    resp,
	})
}
