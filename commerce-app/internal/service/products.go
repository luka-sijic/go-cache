package service

import (
	"app/internal/cache"
	"app/internal/database"
	"app/internal/models"
	"context"
	"fmt"
	"log"
)

func CreateProduct(p *models.Product) bool {
	_, err := database.DB.Exec(context.Background(), `INSERT INTO products (name, price) 
	VALUES ($1,$2)`, p.Name, p.Price)
	if err != nil {
		log.Println("Failed to create product: ", err)
		return false
	}

	return true
}

func GetProduct(id string) *models.Product {
	res := cache.GetProduct(id)
	if res != nil {
		return res
	}
	var product models.Product
	err := database.DB.QueryRow(context.Background(), `SELECT id, name, price FROM products WHERE 
		id=$1`, id).Scan(&product.ID, &product.Name, &product.Price)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	cache.SetProduct(product)
	return &product
}

func GetProducts(page, limit int) []models.Product {
	offset := (page - 1) * limit

	var products []models.Product

	rows, err := database.DB.Query(context.Background(), "SELECT id, name, price FROM products LIMIT $1 OFFSET $2", limit, offset)
	if err != nil {
		log.Println(err)
		return nil
	}
	defer rows.Close()

	for rows.Next() {
		var product models.Product
		err := rows.Scan(&product.ID, &product.Name, &product.Price)
		if err != nil {
			log.Println(err)
			return nil
		}
		products = append(products, product)
	}

	return products
}

func UpdateProduct(p *models.Product) bool {
	_, err := database.DB.Exec(context.Background(), "UPDATE products SET price=$1 WHERE id=$2", p.Price, p.ID)
	if err != nil {
		log.Println("error: ", err)
		return false
	}
	//id := strconv.Itoa(p.ID)

	//cache.DelProduct(id)
	cache.SetProduct(*p)

	return true
}

