// @Title
// @Description
// @Author  Wangwengang  2023/12/10 12:07
// @Update  Wangwengang  2023/12/10 12:07
package sconfig

var S_CONF Config

type Config struct {
	Slog     Slog            `json:"slog" yaml:"slog" mapstructure:"slog"`
	JWT      JWT             `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Gateway  Gateway         `mapstructure:"gateway" json:"gateway" yaml:"gateway"`
	RPC      RPC             `mapstructure:"rpc" yaml:"rpc" json:"rpc"`
	Nsq      Nsq             `mapstructure:"nsq" yaml:"nsq" json:"nsq"`
	Redis    Redis           `mapstructure:"redis" yaml:"redis" json:"redis"`
	RootPath string          `yaml:"root-path" json:"rootPath" mapstructure:"root-path"`
	DBList   []SpecializedDB `mapstructure:"db-list" json:"db-list" yaml:"db-list"`
}
