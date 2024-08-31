package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gokberkotlu/auto-messaging/entity"
	"github.com/gokberkotlu/auto-messaging/service"
)

type IMessageController interface {
	SwitchAutoMessagingStatus(ctx *gin.Context)
	GetSentMessages() []entity.Message
}

type messageController struct {
	service service.IMessageService
}

func NewMessageController() IMessageController {
	return &messageController{
		service: service.NewMessageService(),
	}
}

func (controller *messageController) SwitchAutoMessagingStatus(ctx *gin.Context) {}

func (controller *messageController) GetSentMessages() []entity.Message {
	return controller.service.GetSentMessages()
}
