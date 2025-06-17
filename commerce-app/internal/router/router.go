package router

import (
	"net/http"
	"github.com/labstack/echo/v4"
	"app/internal/handler"
)

func Router(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "TEST")
	})
	e.GET("/products/:id", handler.GetProduct)
	e.POST("/products", handler.CreateProduct)
	e.GET("/products", handler.GetProducts)
	e.PATCH("/products", handler.UpdateProduct)
	e.GET("/events", handler.StreamEvent)
}
