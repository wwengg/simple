// @Title
// @Description
// @Author  Wangwengang  2023/12/10 12:06
// @Update  Wangwengang  2023/12/10 12:06
package sconfig

type Gateway struct {
	Env                 string `mapstructure:"env" json:"env" yaml:"env"` // 环境值
	PrivateRouterPrefix string `mapstructure:"private-router-prefix" json:"privateRouterPrefix" yaml:"private-router-prefix"`
	PublicRouterPrefix  string `mapstructure:"public-router-prefix" yaml:"public-router-prefix" json:"publicRouterPrefix"`
	Addr                string `mapstructure:"addr" json:"addr" yaml:"addr"` // 端口值

	LimitCountIP int `mapstructure:"iplimit-count" json:"iplimit-count" yaml:"iplimit-count"`
	LimitTimeIP  int `mapstructure:"iplimit-time" json:"iplimit-time" yaml:"iplimit-time"`
}
