// @Title
// @Description
// @Author  Wangwengang  2023/12/13 20:44
// @Update  Wangwengang  2023/12/13 20:44
package sconfig

type Nsq struct {
	Nsqd1Addr      string `json:"nsqd1Addr" yaml:"nsqd1-Addr" mapstructure:"nsqd1-addr"`
	Nsqd2Addr      string `json:"nsqd2Addr" yaml:"nsqd2-Addr" mapstructure:"nsqd2-addr"`
	NsqlookupdAddr string `json:"nsqlookupdAddr" yaml:"nsqlookupd-addr" mapstructure:"nsqlookupd-addr"`
}
