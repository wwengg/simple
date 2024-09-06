package internal

import (
	"github.com/getsentry/sentry-go"
	"go.uber.org/zap/zapcore"
	"time"
)

type Sentry struct {
	level zapcore.Level
}

func (c *Sentry) Write(bytes []byte) (n int, err error) {
	defer sentry.Flush(2 * time.Second)
	switch c.level {
	case zapcore.DebugLevel:
		sentry.CaptureMessage(string(bytes))
		break
	case zapcore.InfoLevel:
		sentry.CaptureMessage(string(bytes))
		break
	case zapcore.WarnLevel:
		sentry.CaptureMessage(string(bytes))
		break
	case zapcore.ErrorLevel:
		sentry.CaptureMessage(string(bytes))
		break
	case zapcore.DPanicLevel:
		sentry.CaptureMessage(string(bytes))
		break
	case zapcore.PanicLevel:
		sentry.CaptureMessage(string(bytes))
		break
	case zapcore.FatalLevel:
		sentry.CaptureMessage(string(bytes))
		break
	default:
		sentry.CaptureMessage(string(bytes))
	}
	return len(bytes), nil
}
