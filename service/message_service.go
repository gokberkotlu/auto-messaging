package service

import (
	"github.com/gokberkotlu/auto-messaging/entity"
	"github.com/gokberkotlu/auto-messaging/repository"
)

type IMessageService interface {
	GetNextTwoUnsentMessages() ([]entity.Message, error)
	GetSentMessages() ([]entity.Message, error)
	UpdateMessageStatusAsSent(message entity.Message) error
}

type messageService struct {
	repository repository.IMessageRepository
}

func NewMessageService() IMessageService {
	messageRepository := repository.NewMessageRepository()

	return &messageService{
		repository: messageRepository,
	}
}

func (service *messageService) GetNextTwoUnsentMessages() ([]entity.Message, error) {
	return service.repository.GetNextTwoUnsentMessages()
}

func (service *messageService) GetSentMessages() ([]entity.Message, error) {
	return service.repository.GetSentMessages()
}

func (service *messageService) UpdateMessageStatusAsSent(message entity.Message) error {
	return service.repository.UpdateMessageStatusAsSent(message)
}
