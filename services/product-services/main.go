package main

import (
	"fmt"
	"log"
	"product-services/api"
	"product-services/application"
	"product-services/config"
	"product-services/database"
)

func main() {
	config := config.New()
	dsn := fmt.Sprintf("file:%s?cache=shared&mode=rwc", config.SqLiteDbPath)
	db, err := database.SetupDatabaseConnection(dsn)
	if err != nil {
		log.Fatalf("Database setup error: %v", err)
	}
	defer db.Close()

	app, err := application.New("sqlite", dsn)
	if err != nil {
		fmt.Println(err)
		return
	}

	apiService := api.New("0.0.0.0:"+config.HttpListenPort, app)
	apiService.Start()

	resp, err := app.GetAllProducts()
	if err != nil {
		log.Fatalf("Failed to fetch products: %v", err)
	}
	if len(resp.Products) == 0 {
		log.Println("Product table is empty, seeding initial data...")
		err := app.SeedInitialProducts()
		if err != nil {
			log.Fatalf("Failed to seed products: %v", err)
		}
	}
	defer apiService.Stop()
	select {}
}
