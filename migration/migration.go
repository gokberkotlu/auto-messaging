package migration

import (
	"fmt"

	"github.com/gokberkotlu/auto-messaging/database"
	"github.com/gokberkotlu/auto-messaging/entity"
	"github.com/gokberkotlu/auto-messaging/repository"
	"gorm.io/gorm"
)

var messageRepository repository.IMessageRepository
var migrationProcessRepository repository.IMigrationProcessRepository
var migrationProcessItems []entity.MigrationProcessItem

func AutoMigrate() {
	initializeMigrationVariables()
	migrateMessageTable()
	migrateMigrationProcessTable()
	processNotOperatedMigrations()
}

func initializeMigrationVariables() {
	messageRepository = repository.NewMessageRepository()
	migrationProcessRepository = repository.NewMigrationProcessRepository()
	migrationProcessItems = []entity.MigrationProcessItem{
		entity.MigrationProcessItem{
			Status: entity.MessageBatchLoad,
			Action: messageRepository.BatchLoad,
		},
	}
}

func processNotOperatedMigrations() {
	if migrationProcess, err := repository.NewMigrationProcessRepository().GetFirst(); err != nil {
		if err == gorm.ErrRecordNotFound {
			// do process actions
			for _, processItem := range migrationProcessItems {
				processItem.Action()
			}

			// insert initial migration process data
			migrationProcessRepository.Create(entity.LastStatus)
		}
	} else {
		for _, processItem := range getUnprocessedItemsSlice(migrationProcess.Status) {
			processItem.Action()
		}
		migrationProcessRepository.Update(entity.LastStatus)
	}
}

func getUnprocessedItemsSlice(status entity.MigrationProcessStatusType) []entity.MigrationProcessItem {
	var result = []entity.MigrationProcessItem{}
	if status < entity.MigrationProcessStatusType(len(migrationProcessItems)-1) {
		result = migrationProcessItems[status+1:]
	}

	return result
}

func migrateMessageTable() {
	db, err := database.GetDB()
	if err != nil {
		fmt.Printf("failed to migrate messages table: %s\n", err)
		return
	}
	db.AutoMigrate(&entity.Message{})
}

func migrateMigrationProcessTable() {
	db, err := database.GetDB()
	if err != nil {
		fmt.Printf("failed to migrate migration process table: %s\n", err)
		return
	}
	db.AutoMigrate(&entity.MigrationProcess{})
}
