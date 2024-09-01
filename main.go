package main

import (
	"log"

	automessager "github.com/gokberkotlu/auto-messaging/auto-messager"
	"github.com/gokberkotlu/auto-messaging/docs"
	"github.com/gokberkotlu/auto-messaging/migration"
	"github.com/gokberkotlu/auto-messaging/server"
	"github.com/joho/godotenv"
)

// @title           Auto Messaging API
// @version         1.0
// @description     "Auto Messaging Application Web Server."
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	// programmatically set swagger info
	docs.SwaggerInfo.Title = "Auto Messaging API"
	docs.SwaggerInfo.Description = "Auto Messaging Application Web Server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http"}

	// load environments
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// auto migration check
	migration.AutoMigrate()

	automessager.Init()

	server.Init()
}
