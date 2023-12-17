// @Title
// @Description
// @Author  Wangwengang  2023/12/17 01:01
// @Update  Wangwengang  2023/12/17 01:01
package global

import "github.com/wwengg/simple/core/sconfig"

type Config struct {
	sconfig.Config `yaml:",inline" mapstructure:",squash"`
	PayKey         string `yaml:"pay-key" yaml:"pay-key"`
}
