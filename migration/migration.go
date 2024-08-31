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
		fmt.Errorf("failed to migrate messages table: %w", err)
	}
	db.AutoMigrate(&entity.Message{})
}
