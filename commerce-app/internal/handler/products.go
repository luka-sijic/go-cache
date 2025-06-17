package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"app/internal/database"
	"app/internal/models"
	"app/internal/service"

	"github.com/labstack/echo/v4"
)


func CreateProduct(c echo.Context) error {
	product := new(models.Product)
	if err := c.Bind(product); err != nil {
        log.Println(err)
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
    }
	res := service.CreateProduct(product)
	if !res {
		return c.JSON(http.StatusInternalServerError, "Failed to create product")
	}

	return c.JSON(http.StatusOK, "Product was created")
}

func GetProduct(c echo.Context) error {
	id := c.Param("id")

	product := service.GetProduct(id)
	if product == nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "No product found"})
	}

	return c.JSON(http.StatusOK, product) 
}

func GetProducts(c echo.Context) error {
	pageParam := c.QueryParam("page")
	limitParam := c.QueryParam("limit")

	var products []models.Product

	page, err := strconv.Atoi(pageParam)
	if err != nil || page < 1 {
		page = 1
	}
	limit, err := strconv.Atoi(limitParam)
	if err != nil || limit < 1 {
		limit = 10
	}

	products = service.GetProducts(page, limit)
	if products == nil {
		log.Println("Failed to ret products")
		return c.JSON(http.StatusInternalServerError, "Failed to ret products")
	}

	return c.JSON(http.StatusOK, products)
}

func UpdateProduct(c echo.Context) error {
	product := new(models.Product)
	if err := c.Bind(product); err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	res := service.UpdateProduct(product)
	if !res {
		return c.JSON(http.StatusInternalServerError, "Failed to update product")
	}
    payload, _ := json.Marshal(product)	
	if err := database.RDB.Publish(context.Background(), "product-updates", payload).Err(); err != nil {
		log.Println("redis.Publish error:", err)
	}

	return c.JSON(http.StatusOK, "Product was updated")
}
	

