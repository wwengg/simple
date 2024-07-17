package sbus

import (
	"github.com/nsqio/go-nsq"
	"time"
)

type NsqConsumer struct {
	topic         string
	channel       string
	nsqLookupAddr string
	concurrency   int
	nsqConsumer   *nsq.Consumer
	handler       nsq.Handler
}

func NewNsqConsumer(topic, channel, nsqlookupAddr string, concurrency, maxInFlight int) (*NsqConsumer, error) {
	nsqConsumer := &NsqConsumer{
		topic:         topic,
		channel:       channel,
		concurrency:   concurrency,
		nsqLookupAddr: nsqlookupAddr,
		nsqConsumer:   nil,
		handler:       NewNsqHandler(topic),
	}
	cfg := nsq.NewConfig()
	cfg.LookupdPollInterval = time.Second
	cfg.LookupdPollTimeout = time.Millisecond * 25
	cfg.MaxInFlight = maxInFlight
	if cfg.MaxInFlight <= 0 {
		cfg.MaxInFlight = 100
	}
	if c, err := nsq.NewConsumer(topic, channel, cfg); err != nil {
		return nil, err
	} else {
		nsqConsumer.nsqConsumer = c
		c.AddConcurrentHandlers(nsqConsumer.handler, concurrency)
		return nsqConsumer, nil
	}
}

func (c *NsqConsumer) Start() error {
	return c.nsqConsumer.ConnectToNSQLookupd(c.nsqLookupAddr)
}
