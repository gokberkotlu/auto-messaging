package main

import (
	"fmt"
	"log"

	automessager "github.com/gokberkotlu/auto-messaging/auto-messager"
	"github.com/gokberkotlu/auto-messaging/migration"
	"github.com/gokberkotlu/auto-messaging/server"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("starting app...")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	migration.AutoMigrate()
	// batchload.ReadCSV()
	automessager.Init()

	server.Init()
}
