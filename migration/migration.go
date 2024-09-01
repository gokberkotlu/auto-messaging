package migration

import (
	"fmt"

	"github.com/gokberkotlu/auto-messaging/database"
	"github.com/gokberkotlu/auto-messaging/entity"
)

func AutoMigrate() {
	migrateMessageTable()
}

func migrateMessageTable() {
	db, err := database.GetDB()
	if err != nil {
		fmt.Printf("failed to migrate messages table: %s\n", err)
		return
	}
	db.AutoMigrate(&entity.Message{})
}
