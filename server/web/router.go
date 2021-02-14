package web

import "github.com/labstack/echo"

func SetRoutes(e *echo.Echo) {
	e.GET("/", HelloHandler)
	e.GET("/product", GetProductsHandler)
	e.GET("/product/:id", GetProductByIDHandler)
	e.POST("/product", AddProductHandler)
	e.PUT("/product/:id", UpdateProductHandler)
	e.DELETE("/product/:id", DeleteProductHandler)
}
