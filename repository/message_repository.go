package repository

import (
	"fmt"

	batchload "github.com/gokberkotlu/auto-messaging/batch-load"
	"github.com/gokberkotlu/auto-messaging/database"
	"github.com/gokberkotlu/auto-messaging/entity"
	"gorm.io/gorm"
)

type IMessageRepository interface {
	BatchLoad()
	GetNextTwoUnsentMessages() ([]entity.Message, error)
	GetSentMessages() ([]entity.Message, error)
	UpdateMessageStatusAsSent(message entity.Message) error
}

type MessageRepository struct {
	db *gorm.DB
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

func (repository *MessageRepository) BatchLoad() {
	messages := batchload.ReadCSV()
	err := database.CheckIfDbConnectionInitialized()
	if err != nil {
		return
	}

	result := repository.db.Create(&messages)
	if result.Error != nil {
		fmt.Println("Error inserting batch data:", result.Error)
	} else {
		fmt.Printf("Inserted %d rows successfully.\n", result.RowsAffected)
	}
}
func (repository *MessageRepository) GetNextTwoUnsentMessages() ([]entity.Message, error) {
	err := database.CheckIfDbConnectionInitialized()
	if err != nil {
		return nil, err
	}

	var messages []entity.Message

	if err := repository.db.Where("status = ?", entity.Active).Order("id").Limit(2).Find(&messages).Error; err != nil {
		return nil, fmt.Errorf(`"get next to unsent messages" query failed: %v`, err)
	}

	return messages, nil
}

func (repository *MessageRepository) GetSentMessages() ([]entity.Message, error) {
	err := database.CheckIfDbConnectionInitialized()
	if err != nil {
		return nil, err
	}

	var messages []entity.Message

	if err := repository.db.Where("status = ?", entity.Sent).Order("id").Find(&messages).Error; err != nil {
		return nil, fmt.Errorf(`"get sent messages" query failed: %s`, err.Error())
	}

	return messages, nil
}

func (repository *MessageRepository) UpdateMessageStatusAsSent(message entity.Message) error {
	if err := repository.db.Model(&message).Update("status", entity.Sent).Error; err != nil {
		return fmt.Errorf(`"update message status as sent" query failed: %s`, err.Error())
	}

	return nil
}
