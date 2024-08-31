package service

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	automessager "github.com/gokberkotlu/auto-messaging/auto-messager"
	"github.com/gokberkotlu/auto-messaging/entity"
	"github.com/gokberkotlu/auto-messaging/repository"
)

type IMessageService interface {
	GetNextTwoUnsentMessages() []entity.Message
	GetSentMessages() []entity.Message
	UpdateMessageStatusAsSent(message entity.Message)
	SwitchAutoMessagingMode(ctx *gin.Context)
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

func (service *messageService) UpdateMessageStatusAsSent(message entity.Message) {
	service.repository.UpdateMessageStatusAsSent(message)
}

func (service *messageService) SwitchAutoMessagingMode(ctx *gin.Context) {
	activeParam := ctx.Param("active")
	boolActiveParam, err := strconv.ParseBool(activeParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid mode value. Use 'true' or 'false'.",
		})
		return
	}

	autoMessagerInstance := automessager.GetAutoMessager()

	if autoMessagerInstance.GetMode() != boolActiveParam {
		var action string
		if boolActiveParam {
			autoMessagerInstance.RecreateTicker()
			autoMessagerInstance.Start()
			action = "enabled"
		} else {
			autoMessagerInstance.Stop()
			action = "disabled"
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("auto messager %s", action),
		})

		return
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"error": "mode not changed",
		})
	}
}
