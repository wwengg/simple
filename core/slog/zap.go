// @Title
// @Description
// @Author  Wangwengang  2023/12/12 09:06
// @Update  Wangwengang  2023/12/12 09:06
package slog

import (
	"fmt"
	"github.com/wwengg/simple/core/sconfig"
	"github.com/wwengg/simple/core/slog/internal"
	"github.com/wwengg/simple/core/utils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

//type Field = zap.Field

type Zap struct {
	logger *zap.Logger
	sugar  *zap.SugaredLogger
	config *sconfig.Slog
}

func NewZapLog(config *sconfig.Slog) *Zap {
	if config == nil {
		panic(fmt.Errorf("请在config.yaml中配置slog \n"))
	}
	if ok, _ := utils.PathExists(config.Director); !ok { // 判断是否有Director文件夹
		fmt.Printf("create %v directory\n", config.Director)
		_ = os.Mkdir(config.Director, os.ModePerm)
	}

	cores := internal.GetZapCores(config)
	logger := zap.New(zapcore.NewTee(cores...))

	if config.ShowLine {
		logger = logger.WithOptions(zap.AddCaller(), zap.AddCallerSkip(1))
	}
	defer logger.Sync()
	sugar := logger.Sugar()

	return &Zap{
		logger: logger,
		sugar:  sugar,
		config: config,
	}
}

func (z *Zap) Debug(msg string, fields ...Field) {
	z.logger.Debug(msg, fields...)
}

func (z *Zap) Info(msg string, fields ...Field) {
	z.logger.Info(msg, fields...)
}

func (z *Zap) Error(msg string, fields ...Field) {
	z.logger.Error(msg, fields...)
}

func (z *Zap) Warn(msg string, fields ...Field) {
	z.logger.Warn(msg, fields...)
}

func (z *Zap) Infof(format string, a ...interface{}) {
	z.sugar.Infof(format, a)
}

func (z *Zap) Debugf(format string, a ...interface{}) {
	z.sugar.Debugf(format, a)
}

func (z *Zap) Errorf(format string, a ...interface{}) {
	z.sugar.Errorf(format, a)
}

func (z *Zap) Warnf(format string, a ...interface{}) {
	z.sugar.Warnf(format, a)
}
