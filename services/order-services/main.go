package main

import (
	"fmt"
	"order-services/api"
	"order-services/application"
	"order-services/config"
	"order-services/database"

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
	apiService := api.New("0.0.0.0:"+config.HttpListenPort, app)
	apiService.Start()
	defer apiService.Stop()
	select {}
}
