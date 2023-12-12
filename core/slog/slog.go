// @Title
// @Description
// @Author  Wangwengang  2023/12/12 09:05
// @Update  Wangwengang  2023/12/12 09:05
package slog

type Slog interface {
	Debug(msg string, fields ...Field)
	Info(msg string, fields ...Field)
	Warn(msg string, fields ...Field)
	Error(msg string, fields ...Field)
	Infof(format string, a ...any)
	Debugf(format string, a ...any)
	Errorf(format string, a ...any)
	Warnf(format string, a ...any)
}
