package sbus

import (
	"github.com/nsqio/go-nsq"
	"github.com/wwengg/simple/core/slog"
	"time"
)

type NSQ struct {
	Producer *nsq.Producer
	Consumer *nsq.Consumer
}

func NewNSQConsumer(lookupAddr string, maxInFlight, concurrency int, handler nsq.Handler) (*NSQ, error) {
	cfg := nsq.NewConfig()
	cfg.LookupdPollInterval = time.Second
	cfg.LookupdPollTimeout = time.Millisecond * 25
	cfg.MaxInFlight = maxInFlight
	if cfg.MaxInFlight <= 0 {
		cfg.MaxInFlight = 100
	}
	if c, err := nsq.NewConsumer("", "", cfg); err != nil {
		slog.Ins().Errorf("create nsq consumer error, %v", err)
		return nil, err
	} else {
		//c.SetLogger(nil,0)  // 屏蔽系统日志
		c.AddConcurrentHandlers(handler, concurrency)
	}
}
