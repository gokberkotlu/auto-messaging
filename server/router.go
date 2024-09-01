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
	v1 := r.Group("/api/v1")
	{
		messages := v1.Group("message")
		{
			messages.GET("/switch-auto-messaging-mode/:active", messageController.SwitchAutoMessagingMode)

			messages.GET("/get-sent-messages", messageController.GetSentMessages)
		}

	}

	return r
}
