package handler

import (
	"github.com/GMaikerYactayo/crud-api-goland/middleware"
	"github.com/labstack/echo"
)

// RouteProduct configures product-related routes in a given ServeMux,
// using the provided storage to handle product operations.
func RouteProduct(e *echo.Echo, storage Storage) {
	h := newProductService(storage)
	product := e.Group("/v1/products")
	product.Use(middleware.Authentication)
	product.POST("", h.create)
	product.GET("", h.getAll)
	product.GET("/:id", h.getByID)
	product.PUT("/:id", h.update)
	product.DELETE("/:id", h.delete)
}

func RuteUser(e *echo.Echo, userStorage UserStorage) {
	h := newLoginService(userStorage)
	e.POST("/v1/login", h.login)
}
