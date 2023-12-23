// @Title
// @Description
// @Author  Wangwengang  2023/12/16 14:44
// @Update  Wangwengang  2023/12/16 14:44
package sconfig

type RpcService struct {
	ServiceAddr string `json:"serviceAddr" yaml:"service-addr" mapstructure:"service-addr"`
	Port        string `json:"port" yaml:"port" mapstructure:"port"`
}
