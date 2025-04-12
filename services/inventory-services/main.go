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
	rabbitMQURL := "amqp://guest:guest@localhost:5672/"
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
	apiService := api.New(":"+config.HttpListenPort, app)
	apiService.Start()
	defer apiService.Stop()
	select {}
}
