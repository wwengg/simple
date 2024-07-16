// @Title
// @Description
// @Author  Wangwengang  2023/12/13 20:44
// @Update  Wangwengang  2023/12/13 20:44
package sconfig

type Nsq struct {
	NsqlookupdAddr string   `json:"nsqlookupdAddr" yaml:"nsqlookupd-addr" mapstructure:"nsqlookupd-addr"`
	NsqdAddrList   []string `json:"nsqdAddrList" yaml:"nsqd-addr-list" mapstructure:"nsqd-addr-list"`
}
