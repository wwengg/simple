// @Title
// @Description
// @Author  Wangwengang  2023/12/17 11:55
// @Update  Wangwengang  2023/12/17 11:55
package global

import "github.com/wwengg/simple/core/sconfig"

type Config struct {
	sconfig.Config `yaml:",inline" mapstructure:",squash"`
	// 下方加入自己需要额外加入的配置
}
