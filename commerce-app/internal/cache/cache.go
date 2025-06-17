package cache

import (
	"context"
	"app/internal/models"
	"strconv"
	"fmt"
	"log"
	"app/internal/database"
)

func GetProduct(id string) *models.Product {
	key := fmt.Sprintf("products:%s", id)
	fields, err := database.RDB.HGetAll(context.Background(), key).Result()
	if err != nil {
		log.Println("Product not found", err)
		return nil
	}
	pID, err := strconv.Atoi(id)
	if err != nil {
		return nil
	}

	priceStr, err := strconv.ParseFloat(fields["price"], 64)
	if err != nil {
		return nil
	}

	product := &models.Product{
		ID: pID,
		Name: fields["name"],
		Price: priceStr,
	}

	return product
}

func SetProduct(p models.Product) {
	id := strconv.Itoa(p.ID)
	key := fmt.Sprintf("products:%s", id)
	fields := map[string]interface{} {
		"name": p.Name,
		"price": p.Price,
	}
	_ = database.RDB.HSet(context.Background(), key, fields)
}

func DelProduct(id string) bool {
	 key := fmt.Sprintf("products:%s", id)
	 _ = database.RDB.Del(context.Background(), key)

	 return true
 }
