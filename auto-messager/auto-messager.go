package automessager

import (
	"fmt"
	"sync"
	"time"
)

type IAutoMessager interface {
	Start()
	Stop()
	RecreateTicker()
	GetMode() bool
}

type AutoMessager struct {
	Ticker *time.Ticker
	QuitCh chan struct{}
	Mode   bool
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
		Ticker: getTicker(),
		QuitCh: make(chan struct{}),
		Mode:   true,
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

func Init() {
	autoMessagerInstance := GetAutoMessager()
	autoMessagerInstance.Start()
}

func (autoMessager *AutoMessager) GetMode() bool {
	return autoMessager.Mode
}
