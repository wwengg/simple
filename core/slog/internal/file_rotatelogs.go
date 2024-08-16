// @Title
// @Description
// @Author  Wangwengang  2023/12/12 09:38
// @Update  Wangwengang  2023/12/12 09:38
package internal

import (
	"github.com/wwengg/simple/core/sconfig"
	"go.uber.org/zap/zapcore"
	"os"
)

var FileRotatelogs = new(fileRotatelogs)

type fileRotatelogs struct{}

// GetWriteSyncer 获取 zapcore.WriteSyncer

func (r *fileRotatelogs) GetWriteSyncer(level string, config *sconfig.Slog) zapcore.WriteSyncer {
	fileWriter := NewCutter(config.Director, level, config.IsAllInOne, WithCutterFormat("2006-01-02"))
	if config.LogInConsole {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(fileWriter))
	}
	return zapcore.AddSync(fileWriter)
}
