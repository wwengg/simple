// @Title
// @Description
// @Author  Wangwengang  2023/12/12 09:08
// @Update  Wangwengang  2023/12/12 09:08
package sconfig

type Slog struct {
	Level         string `mapstructure:"level" json:"level" yaml:"level"`                            // 级别
	Format        string `mapstructure:"format" json:"format" yaml:"format"`                         // 输出
	Director      string `mapstructure:"director" json:"director"  yaml:"director"`                  // 日志文件夹
	EncodeLevel   string `mapstructure:"encode-level" json:"encode-level" yaml:"encode-level"`       // 编码级
	StacktraceKey string `mapstructure:"stacktrace-key" json:"stacktrace-key" yaml:"stacktrace-key"` // 栈名
	Prefix        string `mapstructure:"prefix" json:"prefix" yaml:"prefix"`                         // 前缀

	MaxAge       int  `mapstructure:"max-age" json:"max-age" yaml:"max-age"`                      // 日志留存时间
	ShowLine     bool `mapstructure:"show-line" json:"show-line" yaml:"show-line"`                // 显示行
	LogInConsole bool `mapstructure:"log-in-console" json:"log-in-console" yaml:"log-in-console"` // 输出控制台
	LogInNsq     bool `mapstructure:"log-in-nsq" json:"log-in-nsq" yaml:"log-in-nsq"`             // 输出到nsq
}
