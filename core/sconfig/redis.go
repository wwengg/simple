// @Title
// @Description
// @Author  Wangwengang  2023/12/14 04:30
// @Update  Wangwengang  2023/12/14 04:30
package sconfig

type Redis struct {
	Addr     string `mapstructure:"addr" json:"addr" yaml:"addr"`             // 服务器地址:端口
	Password string `mapstructure:"password" json:"password" yaml:"password"` // 密码
	DB       int    `mapstructure:"db" json:"db" yaml:"db"`                   // redis的哪个数据库
}
