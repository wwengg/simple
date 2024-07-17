package sbus

import "github.com/nsqio/go-nsq"

type NsqHandler struct {
	Topic string
}

func NewNsqHandler(topic string) *NsqHandler {
	return &NsqHandler{
		Topic: topic,
	}
}

func (a *NsqHandler) HandleMessage(message *nsq.Message) error {
	return nil
}
