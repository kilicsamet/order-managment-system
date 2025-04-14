package main

import (
	"fmt"
	"inventory-service/api"
	"inventory-service/application"
	"inventory-service/config"
	"inventory-service/database"
	"log"
)

func main() {
	config := config.New()
	dsn := fmt.Sprintf("file:%s?cache=shared&mode=rwc", config.SqLiteDbPath)
	rabbitMQURL := fmt.Sprintf(
		"amqp://%s:%s@%s:%s/",
		config.RabbitMQUser,
		config.RabbitMQPass,
		config.RabbitMQHost,
		config.RabbitMQPort,
	)
	db, err := database.SetupDatabaseConnection(dsn)
	if err != nil {
		log.Fatalf("Database setup error: %v", err)
	}
	defer db.Close()

	app, err := application.New("sqlite", dsn, rabbitMQURL)
	if err != nil {
		fmt.Println(err)
		return
	}
	go app.ListenToOrderQueue()
	apiService := api.New("0.0.0.0:"+config.HttpListenPort, app)
	apiService.Start()

	resp, err := app.GetAllInventoryProducts()
	if err != nil {
		log.Fatalf("Failed to fetch products: %v", err)
	}
	if len(resp.InventoryProducts) == 0 {
		log.Println("Product table is empty, seeding initial data...")
		err := app.SeedInitialInventoryProducts()
		if err != nil {
			log.Fatalf("Failed to seed products: %v", err)
		}
	}

	defer apiService.Stop()
	select {}
}
