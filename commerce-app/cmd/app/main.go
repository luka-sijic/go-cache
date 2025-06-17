package main

import (
	"app/internal/server"
	"app/internal/database"
	"app/internal/cache"
)

func main() {
	database.Connect()
	go cache.SubscribeUpdates()
	defer database.Close()

	server.Start()
}
