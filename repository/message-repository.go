package repository

import (
	"fmt"
	"log"

	"github.com/gokberkotlu/auto-messaging/database"
	"github.com/gokberkotlu/auto-messaging/entity"
	"gorm.io/gorm"
)

type MessageRepository struct {
	db *gorm.DB
}

type IMessageRepository interface {
	BatchLoad(messages []entity.Message)
	GetNextTwoUnsentMessages() ([]entity.Message, error)
	GetSentMessages() ([]entity.Message, error)
	UpdateMessageStatusAsSent(message entity.Message) error
	checkIfDbConnectionExists() error
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

func (repository *MessageRepository) BatchLoad(messages []entity.Message) {
	err := repository.checkIfDbConnectionExists()
	if err != nil {
		fmt.Printf(err.Error())
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
	err := repository.checkIfDbConnectionExists()
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
	err := repository.checkIfDbConnectionExists()
	if err != nil {
		return nil, err
	}

	var messages []entity.Message

	if err := repository.db.Where("status = ?", entity.Sent).Order("id").Find(&messages).Error; err != nil {
		log.Fatalf(`"get sent messages" query failed: %v`, err)
	}

	return messages, nil
}

func (repository *MessageRepository) UpdateMessageStatusAsSent(message entity.Message) error {
	if err := repository.db.Model(&message).Update("status", entity.Sent).Error; err != nil {
		return fmt.Errorf(`"update message status as sent" query failed: %v`, err)
	}

	return nil
}

func (repository *MessageRepository) checkIfDbConnectionExists() error {
	if repository == nil || (repository != nil && repository.db == nil) {
		return fmt.Errorf("database connection is not initialized")
	}

	return nil
}
