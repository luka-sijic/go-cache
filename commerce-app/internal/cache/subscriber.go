package cache

import (
	"app/internal/database"
	"log"
	"app/internal/broker"
	"app/internal/models"
	"context"
	"encoding/json"
)

func SubscribeUpdates() {
	pubsub := database.RDB.Subscribe(context.Background(), "product-updates")
	log.Println("REDIS")
	defer pubsub.Close()

	ch := pubsub.Channel()
	for msg := range ch {
		var p models.Product
		if err := json.Unmarshal([]byte(msg.Payload), &p); err != nil {
			log.Println("unmarshal error:", err)
			continue
		}

		broker.Brokerk.Broadcast(msg.Payload)
	}
}
