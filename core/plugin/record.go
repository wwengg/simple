package plugin

import (
	"context"
	"fmt"

	"github.com/wwengg/simple/core/slog"
	"go.uber.org/zap"
)

type RecordPlugin struct {
	logger slog.Slog
}

func NewRecordPlugin(slog slog.Slog) *RecordPlugin {
	return &RecordPlugin{
		logger: slog,
	}
}

func (p *RecordPlugin) PreCall(ctx context.Context, serviceName, methodName string, args interface{}) (interface{}, error) {
	p.logger.Info(fmt.Sprintf("Start[%s][%s]", serviceName, methodName), zap.Any("args", args))
	return args, nil
}

func (p *RecordPlugin) PostCall(ctx context.Context, serviceName, methodName string, args, reply interface{}, err error) (interface{}, error) {
	p.logger.Info(fmt.Sprintf("Finish[%s][%s]", serviceName, methodName), zap.Any("reply", reply))
	return reply, nil
}
