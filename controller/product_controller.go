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
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}

	product := entities.Products{
		NamaProduct: req.NamaProduct,
		Harga:       req.Harga,
		Quantity:    req.Quantity,
		IdBrand:     req.IdBrand,
	}

	if err := config.DB.Create(&product).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create product"})
	}

	return c.JSON(http.StatusCreated, response.CreateProductResponse{ID: product.ID})
}

func GetProductById(c echo.Context) error {
	id := c.Param("id")

	productID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid product ID"})
	}

	var product entities.Products
	if err := config.DB.First(&product, productID).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Product not found"})
	}

	resp := response.GetProductByIdResponse{
		ID:          product.ID,
		NamaProduct: product.NamaProduct,
		Harga:       product.Harga,
		Quantity:    product.Quantity,
		IdBrand:     product.IdBrand,
	}

	return c.JSON(http.StatusOK, resp)
}

func UpdateProduct(c echo.Context) error {
	id := c.Param("id")

	productID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid product ID"})
	}

	var req request.UpdateProductRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}

	var product entities.Products
	if err := config.DB.First(&product, productID).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Product not found"})
	}

	product.NamaProduct = req.NamaProduct
	product.Harga = req.Harga
	product.Quantity = req.Quantity
	product.IdBrand = req.IdBrand

	if err := config.DB.Save(&product).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update product"})
	}

	resp := response.UpdateProductResponse{
		NamaProduct: product.NamaProduct,
		Harga:       product.Harga,
		Quantity:    product.Quantity,
		IdBrand:     product.IdBrand,
	}
	return c.JSON(http.StatusOK, resp)
}

func DeleteProduct(c echo.Context) error {
	id := c.Param("id")

	productID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid product ID"})
	}

	var product entities.Products
	if err := config.DB.First(&product, productID).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Product not found"})
	}

	if err := config.DB.Delete(&product).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to delete product"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Product deleted successfully"})
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
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to count products"})
	}

	var products []entities.Products
	if err := config.DB.Offset(offset).Limit(sizeInt).Find(&products).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch products"})
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

	return c.JSON(http.StatusOK, response.GetAllProductsResponse{
		PageInformation: pageInfo,
		ListData:        productList,
	})
}
