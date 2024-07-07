package slog

import (
	"context"
	"time"

	"gorm.io/gorm/logger"
)

type GormZapLogger struct {
	Logger Slog
}

//	type Interface interface {
//		LogMode(LogLevel) Interface
//		Info(context.Context, string, ...interface{})
//		Warn(context.Context, string, ...interface{})
//		Error(context.Context, string, ...interface{})
//		Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error)
//	}
func NewGormZapLogger(slog Slog) logger.Interface {
	return &GormZapLogger{
		Logger: slog,
	}
}

func (g *GormZapLogger) LogMode(logger.LogLevel) logger.Interface {
	return g
}

func (g *GormZapLogger) Info(ctx context.Context, message string, data ...interface{}) {
	g.Logger.Infof(message, data...)
}

func (g *GormZapLogger) Warn(ctx context.Context, message string, data ...interface{}) {
	g.Logger.Warnf(message, data...)
}

func (g *GormZapLogger) Error(ctx context.Context, message string, data ...interface{}) {
	g.Logger.Errorf(message, data...)
}

func (g *GormZapLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	// If you want to log SQL traces, you can implement this function.
}
