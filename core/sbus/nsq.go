package sbus

import (
	"fmt"
	"github.com/nsqio/go-nsq"
	"github.com/wwengg/simple/core/sconfig"
	"github.com/wwengg/simple/core/slog"
	"time"
)

type NSQ struct {
	ProducerList []*nsq.Producer
	Consumer     *nsq.Consumer
	lookupAddr   string
}

func NewNSQConsumer(lookupAddr string, maxInFlight, concurrency int, handler nsq.Handler, nsqConf *sconfig.Nsq) (*NSQ, error) {
	cfg := nsq.NewConfig()
	cfg.LookupdPollInterval = time.Second
	cfg.LookupdPollTimeout = time.Millisecond * 25
	cfg.MaxInFlight = maxInFlight
	if cfg.MaxInFlight <= 0 {
		cfg.MaxInFlight = 100
	}
	var pList []*nsq.Producer
	c, err := nsq.NewConsumer("", "", cfg)
	if err != nil {
		slog.Ins().Errorf("create nsq consumer error, %v", err)
		return nil, err
	}
	slog.Ins().Info("create nsq consumer success")
	for _, nsqdAddr := range nsqConf.NsqdAddrList {
		p, err := nsq.NewProducer(nsqdAddr, cfg)
		if err == nil {
			slog.Ins().Infof("create nsq producer success addr: %s", nsqdAddr)
			pList = append(pList, p)
		}
	}
	if len(pList) == 0 {
		return nil, fmt.Errorf("create NSQ error producer nil")
	}
	//c.SetLogger(nil,0)  // 屏蔽系统日志
	//c.AddConcurrentHandlers(handler, concurrency)
	return &NSQ{
		ProducerList: pList,
		Consumer:     c,
		lookupAddr:   lookupAddr,
	}, nil
}

func (n *NSQ) Start() {

}

func (n *NSQ) Stop() {}
