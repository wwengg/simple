package sbus

import (
	"github.com/nsqio/go-nsq"
	"github.com/wwengg/simple/core/sbus/sface"
	"github.com/wwengg/simple/core/slog"
	"github.com/wwengg/simple/core/snet"
	"strconv"
	"time"
)

type TopicEnum int32

var TopicEnum_name = map[int32]string{}

var TopicEnum_value = map[string]int32{}

func (t TopicEnum) String() string {
	return enumName(TopicEnum_name, int32(t))
}

func enumName(m map[int32]string, v int32) string {
	s, ok := m[v]
	if ok {
		return s
	}
	return strconv.Itoa(int(v))
}

func SetTopicEnumName(v map[int32]string) {
	TopicEnum_name = v
}

func SetTopicEnumValue(v map[string]int32) {
	TopicEnum_value = v
}

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

func (c *NsqConsumer) StartReader() error {
	return c.nsqConsumer.ConnectToNSQLookupd(c.nsqLookupAddr)
}

func (c *NsqConsumer) Stop() {
	c.nsqConsumer.Stop()
}

type NsqProducer struct {
}

type Nsq struct {
	snet.BaseConnection
	producers []*nsq.Producer
	Consumers []*NsqConsumer

	// The message management module that manages MsgID and the corresponding processing method
	// (消息管理MsgID和对应处理方法的消息管理模块)
	taskHandler sface.STaskHandler
	// Buffered channel used for message communication between the read and write goroutines
	// (有缓冲管道，用于读、写两个goroutine之间的消息通信)
	msgBuffChan chan []byte
}

func (n *Nsq) Start() {
	defer func() {
		if err := recover(); err != nil {
			slog.Ins().Errorf("Nsq Start() error: %v", err)
		}
	}()
	for _, consumer := range n.Consumers {
		err := consumer.StartReader()
		if err != nil {
			panic(err)
		}
	}
}
