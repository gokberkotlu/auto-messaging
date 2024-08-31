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
	GetNextTwoUnsentMessages() []entity.Message
	GetSentMessages() []entity.Message
	UpdateMessageStatusAsSent(message entity.Message)
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

func (repository *MessageRepository) BatchLoad(messages []entity.Message) {
	result := repository.db.Create(&messages)
	if result.Error != nil {
		fmt.Println("Error inserting batch data:", result.Error)
	} else {
		fmt.Printf("Inserted %d rows successfully.\n", result.RowsAffected)
	}
}
func (repository *MessageRepository) GetNextTwoUnsentMessages() []entity.Message {
	var messages []entity.Message

	if err := repository.db.Where("status = ?", entity.Active).Order("id").Limit(2).Find(&messages).Error; err != nil {
		log.Fatalf(`"get next to unsent messages" query failed: %v`, err)
	}

	return messages
}

func (repository *MessageRepository) GetSentMessages() []entity.Message {
	var messages []entity.Message

	if err := repository.db.Where("status = ?", entity.Sent).Find(&messages).Error; err != nil {
		log.Fatalf(`"get sent messages" query failed: %v`, err)
	}

	return messages
}

func (repository *MessageRepository) UpdateMessageStatusAsSent(message entity.Message) {
	repository.db.Model(&message).Update("status", entity.Sent)
}

func (repository *MessageRepository) Save() {}

func (repository *MessageRepository) Update() {}
