package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NetRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"test": "success",
		})
	})

	return r
}
