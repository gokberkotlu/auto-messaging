package controller

import (
	"github.com/gin-gonic/gin"
	automessager "github.com/gokberkotlu/auto-messaging/auto-messager"
	"github.com/gokberkotlu/auto-messaging/entity"
	"github.com/gokberkotlu/auto-messaging/service"
)

type IMessageController interface {
	SwitchAutoMessagingMode(ctx *gin.Context)
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

func (controller *messageController) SwitchAutoMessagingMode(ctx *gin.Context) {
	automessager := automessager.GetAutoMessager()
	automessager.Switch(controller.service, ctx)
}

func (controller *messageController) GetSentMessages() []entity.Message {
	return controller.service.GetSentMessages()
}
