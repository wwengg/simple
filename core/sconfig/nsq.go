// @Title
// @Description
// @Author  Wangwengang  2023/12/13 20:44
// @Update  Wangwengang  2023/12/13 20:44
package sconfig

type Nsq struct {
	NsqlookupdAddr    string   `json:"nsqlookupdAddr" yaml:"nsqlookupd-addr" mapstructure:"nsqlookupd-addr"`
	NsqdAddrList      []string `json:"nsqdAddrList" yaml:"nsqd-addr-list" mapstructure:"nsqd-addr-list"`
	WorkerPoolSize    uint32   `json:"workerPoolSize" yaml:"worker-pool-size" mapstructure:"worker-pool-size"`
	MaxTaskChanLen    uint32   `json:"maxTaskChanLen" yaml:"max-task-chan-len" mapstructure:"max-task-chan-len"` // The maximum length of the send buffer message queue.(SendBuffMsg发送消息的缓冲最大长度)
	MaxNsqDataChanLen uint32   `json:"maxNsqDataChanLen" yaml:"max-nsq-data-chan-len" mapstructure:"max-nsq-data-chan-len"`
	Channel           string   `json:"channel" yaml:"channel" mapstructure:"channel"`
	Concurrency       int      `json:"concurrency" yaml:"concurrency" mapstructure:"concurrency"`
	MaxInFlight       int      `json:"maxInFlight" yaml:"max-in-flight" mapstructure:"max-in-flight"`
}
