// @Title
// @Description
// @Author  Wangwengang  2023/12/10 18:45
// @Update  Wangwengang  2023/12/10 18:45
package sconfig

type RPC struct {
	RegisterType string   `json:"registerType" yaml:"register-type" mapstructure:"register-type"`
	RegisterAddr []string `json:"registerAddr" yaml:"register-addr" mapstructure:"register-addr"`
	BasePath     string   `json:"basePath" yaml:"base-path" mapstructure:"base-path"`
}
