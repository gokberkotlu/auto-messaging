package service

import (
	"github.com/gokberkotlu/auto-messaging/entity"
	"github.com/gokberkotlu/auto-messaging/repository"
)

type IMessageService interface {
	GetNextTwoUnsentMessages() []entity.Message
	GetSentMessages() []entity.Message
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

func (service *messageService) GetNextTwoUnsentMessages() []entity.Message {
	return service.repository.GetNextTwoUnsentMessages()
}

func (service *messageService) GetSentMessages() []entity.Message {
	return service.repository.GetSentMessages()
}
