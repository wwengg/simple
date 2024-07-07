package plugin

import (
	"context"
	"reflect"

	"github.com/wwengg/simple/core/slog"
)

type RecordPlugin struct {
	logger slog.Slog
}

func NewRecordPlugin(slog slog.Slog) *RecordPlugin {
	return &RecordPlugin{
		logger: slog,
	}
}

func (p *RecordPlugin) DoPreCall(ctx context.Context, serviceName, methodName string, args interface{}) (interface{}, error) {
	p.logger.Infof("Start[%s][%s], args: %v", serviceName, methodName, reflect.ValueOf(args))
	return args, nil
}

func (p *RecordPlugin) DoPostCall(ctx context.Context, serviceName, methodName string, args, reply interface{}, err error) (interface{}, error) {
	p.logger.Infof("Finish[%s][%s], reply: %v", serviceName, methodName, reflect.ValueOf(reply))
	return reply, nil
}
