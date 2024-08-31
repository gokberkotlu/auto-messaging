package repository

import (
	"fmt"

	"github.com/gokberkotlu/auto-messaging/database"
	"github.com/gokberkotlu/auto-messaging/entity"
	"gorm.io/gorm"
)

type MessageRepository struct {
	db *gorm.DB
}

type IMessageRepository interface {
	BulkLoad(messages []entity.Message)
	Save()
	Update()
}

func NewMessageRepository() IMessageRepository {
	db, err := database.GetDB()
	if err != nil {
		fmt.Errorf("failed to create message repository: %w", err)
	}
	return &MessageRepository{
		db: db,
	}
}

func (repository *MessageRepository) BulkLoad(messages []entity.Message) {
	result := repository.db.Create(&messages)
	if result.Error != nil {
		fmt.Println("Error inserting batch data:", result.Error)
	} else {
		fmt.Printf("Inserted %d rows successfully.\n", result.RowsAffected)
	}
}
func (repository *MessageRepository) Save()   {}
func (repository *MessageRepository) Update() {}
