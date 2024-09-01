package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	automessager "github.com/gokberkotlu/auto-messaging/auto-messager"
	"github.com/gokberkotlu/auto-messaging/dto"
	"github.com/gokberkotlu/auto-messaging/entity"
	"github.com/gokberkotlu/auto-messaging/service"
)

type IMessageController interface {
	SwitchAutoMessagingMode(ctx *gin.Context)
	GetSentMessages(ctx *gin.Context)
}

type messageController struct {
	service service.IMessageService
}

func NewMessageController() IMessageController {
	return &messageController{
		service: service.NewMessageService(),
	}
}

// ListMessage godoc
// @Summary      Start/Stop Messaging
// @Description  Start/Stop automatic message sending
// @Tags         messages
// @Accept       json
// @Produce      json
// @Param        active   path      bool  true  "Active Mode"
// @Success      200  {array}   dto.SuccessResponse[any]
// @Failure      400  {object}  dto.ErrorResponseDTO
// @Router       /api/v1/message/switch-auto-messaging-mode/{active} [get]
func (controller *messageController) SwitchAutoMessagingMode(ctx *gin.Context) {
	automessager := automessager.GetAutoMessager()
	status, res := automessager.Switch(controller.service, ctx)
	ctx.JSON(status, res)
}

// ListMessage godoc
// @Summary      List messages
// @Description  Retrieve a list of sent messages
// @Tags         messages
// @Accept       json
// @Produce      json
// @Success      200  {array}   dto.SuccessResponse[entity.Message]
// @Failure      500  {object}  dto.ErrorResponseDTO
// @Router       /api/v1/message/get-sent-messages [get]
func (controller *messageController) GetSentMessages(ctx *gin.Context) {
	messages, err := controller.service.GetSentMessages()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.ErrorResponseDTO{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.SuccessResponse[entity.Message]{
		Status:  http.StatusOK,
		Data:    messages,
		Message: "sent messages were successfully retrieved",
	})
}
