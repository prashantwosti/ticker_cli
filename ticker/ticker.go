package ticker

import "github.com/prashantwosti/ticker_cli/service"

type ticker struct{}

type Ticker interface {
	Get(symbol string) (string, service.ErrorCode)
}

func (t *ticker) Get(symbol string) (string, service.ErrorCode) {
	return service.GetTicker(symbol)
}

func NewTicker() Ticker {
	return &ticker{}
}
