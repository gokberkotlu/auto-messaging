package automessager

import "time"

type IAutoMessager interface {
	Start()
	Stop()
}

type AutoMessager struct {
	ticker *time.Ticker
	quitCh chan struct{}
}

const messagingTimeInterval = 2 * time.Minute

func New() IAutoMessager {
	return &AutoMessager{
		ticker: time.NewTicker(messagingTimeInterval),
		quitCh: make(chan struct{}),
	}
}

func (autoMessager *AutoMessager) Start() {
	go func() {
		for {
			select {
			case <-autoMessager.ticker.C:
				// do stuff
			case <-autoMessager.quitCh:
				autoMessager.ticker.Stop()
				return
			}
		}
	}()
}

func (autoMessager *AutoMessager) Stop() {
	autoMessager.quitCh <- struct{}{}
}
