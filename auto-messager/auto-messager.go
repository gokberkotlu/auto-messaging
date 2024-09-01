package automessager

import (
	"fmt"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gokberkotlu/auto-messaging/client"
	"github.com/gokberkotlu/auto-messaging/dto"
	"github.com/gokberkotlu/auto-messaging/service"
)

type IAutoMessager interface {
	Start()
	Stop()
	Switch(messageService service.IMessageService, ctx *gin.Context)
	RecreateTicker()
	GetMode() bool
}

type AutoMessager struct {
	Ticker        *time.Ticker
	QuitCh        chan struct{}
	Mode          bool
	messageClient client.IMessageClient
}

var (
	AutoMessagerInstance IAutoMessager
	lock                 = &sync.Mutex{}
)

const messagingTimeInterval = 2 * time.Second

func GetAutoMessager() IAutoMessager {
	if AutoMessagerInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if AutoMessagerInstance == nil {
			AutoMessagerInstance = newAutoMessager()
		}
	}

	return AutoMessagerInstance
}

func newAutoMessager() *AutoMessager {
	return &AutoMessager{
		Ticker:        getTicker(),
		QuitCh:        make(chan struct{}),
		Mode:          true,
		messageClient: client.New(),
	}
}

func (autoMessager *AutoMessager) RecreateTicker() {
	lock.Lock()
	defer lock.Unlock()
	autoMessager.Ticker = getTicker()
}

func getTicker() *time.Ticker {
	return time.NewTicker(messagingTimeInterval)
}

func (autoMessager *AutoMessager) Start() {
	go func() {
		autoMessager.Mode = true
		for {
			select {
			case <-autoMessager.Ticker.C:
				fmt.Println(time.Now().Format(time.RFC1123))
				autoMessager.messageClient.SendNextTwoUnsentMessages()
			case <-autoMessager.QuitCh:
				autoMessager.Ticker.Stop()
				autoMessager.Mode = false
				return
			}
		}
	}()
}

func (autoMessager *AutoMessager) Stop() {
	autoMessager.QuitCh <- struct{}{}
}

func (autoMessager *AutoMessager) Switch(messageService service.IMessageService, ctx *gin.Context) {
	activeParam := ctx.Param("active")
	boolActiveParam, err := strconv.ParseBool(activeParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponseDTO{
			Status: http.StatusBadRequest,
			Error:  "Invalid mode value. Use 'true' or 'false'.",
		})
		return
	}

	autoMessagerInstance := GetAutoMessager()

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

		ctx.JSON(http.StatusOK, dto.SuccessResponse[any]{
			Status:  http.StatusOK,
			Data:    []any{},
			Message: fmt.Sprintf("auto messager %s", action),
		})

		return
	} else {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponseDTO{
			Status: http.StatusBadRequest,
			Error:  "mode not changed",
		})
	}
}

func Init() {
	autoMessagerInstance := GetAutoMessager()
	autoMessagerInstance.Start()
}

func (autoMessager *AutoMessager) GetMode() bool {
	return autoMessager.Mode
}
