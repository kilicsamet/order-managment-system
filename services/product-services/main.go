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
	defer apiService.Stop()
	select {}
}
