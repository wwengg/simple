// @Title
// @Description
// @Author  Wangwengang  2023/12/10 12:07
// @Update  Wangwengang  2023/12/10 12:07
package sconfig

var S_CONF Config

type Config struct {
	Slog    Slog    `json:"slog" yaml:"slog" mapstructure:"slog"`
	JWT     JWT     `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Gateway Gateway `mapstructure:"gateway" json:"gateway" yaml:"gateway"`
	RPC     RPC     `mapstructure:"rpc" yaml:"rpc" json:"rpc"`

	RootPath string `yaml:"root-path" json:"rootPath" mapstructure:"root-path"`
}
