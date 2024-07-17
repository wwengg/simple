package sface

type STask interface {
	GetConsumer() SConsumer
	GetProducer() SProducer
	
	GetData() []byte    // 获取请求消息的数据
	GetTopic() string   // 获取topic
	GetChannel() string // 获取Channel

	GetMsg() SMsg
	GetWorkerID() uint32
}
