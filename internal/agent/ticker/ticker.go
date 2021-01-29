package ticker

import (
	"time"

	"github.com/frauniki/ping-mesh/pkg/config"
)

func NewTicker() *time.Ticker {
	interval, _ := time.ParseDuration(config.Config.Interval)
	return time.NewTicker(interval)
}
