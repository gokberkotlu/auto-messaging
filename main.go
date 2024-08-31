package main

import (
	"log"

	"github.com/gokberkotlu/auto-messaging/migration"
	"github.com/gokberkotlu/auto-messaging/server"
	"github.com/joho/godotenv"
)

func main() { // Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	migration.AutoMigrate()
	// batchload.ReadCSV()
	server.Init()
}
