package sbus

import (
	"github.com/nsqio/go-nsq"
	"github.com/wwengg/simple/core/slog"
	"go.uber.org/zap"
	"runtime"
)

type NsqHandler struct {
	Topic string
}

func NewNsqHandler(topic string) *NsqHandler {
	return &NsqHandler{
		Topic: topic,
	}
}

func (a *NsqHandler) HandleMessage(message *nsq.Message) error {
	defer func() {
		if err := recover(); err != nil {
			slog.Ins().Errorf("Nsq Start() error: %v", err)
			for i := 0; ; i++ {
				pc, file, line, ok := runtime.Caller(i)
				if !ok {
					break
				}
				slog.Ins().Error("NsqHandler", zap.Any("pc", pc), zap.Any("file", file), zap.Any("line", line))
			}
		}
	}()
	defer func() {
		e := recover()
		if e != nil {
			if LOG != nil && task != nil {
				LOG.Errorf("Panic error=%v,Topic=%s,Cmd=%v", e, task.Topic, task.Cmd)

			}
		}
	}()
	return nil
}
