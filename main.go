package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mhdianrush/ecommerce-project/config"
	"github.com/mhdianrush/ecommerce-project/controller"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	config.ConnectDB()

	e.POST("/products", controller.CreateProduct)
	e.GET("/products/:id", controller.GetProductById)
	e.PUT("/products/:id", controller.UpdateProduct)
	e.DELETE("/products/:id", controller.DeleteProduct)
	e.GET("/products", controller.GetAllProducts)

	e.POST("/brands", controller.CreateBrand)
	e.DELETE("/brands/:id", controller.DeleteBrand)

	e.Logger.Fatal(e.Start(":8000"))
}
