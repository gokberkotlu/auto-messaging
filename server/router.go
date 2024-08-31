package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gokberkotlu/auto-messaging/controller"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"test": "success",
		})
	})

	// message endpoints
	var messageController controller.IMessageController = controller.NewMessageController()
	messageGroup := r.Group("message")

	messageGroup.GET("/switch-auto-messaging-mode/:active", messageController.SwitchAutoMessagingMode)

	messageGroup.GET("/get-sent-messages", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, messageController.GetSentMessages())
	})

	return r
}
