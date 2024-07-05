package plugin

import (
	"context"
	"fmt"

	"github.com/opentracing/opentracing-go"
	"github.com/smallnest/rpcx/protocol"
)

type JaegerPlugin struct {
	span opentracing.Span
}

func NewJaegerPlugin(t opentracing.Tracer) *JaegerPlugin {
	opentracing.SetGlobalTracer(t)
	return &JaegerPlugin{}
}

func (p *JaegerPlugin) DoPreCall(ctx context.Context, serviceName, methodName string, args interface{}) (interface{}, error) {

	return nil, nil
}

func (p *JaegerPlugin) PostReadRequest(ctx context.Context, r *protocol.Message, e error) error {
	sp := r.ServicePath
	sm := r.ServiceMethod

	if sp == "" {
		return nil
	}
	if span, _, err := GenerateSpanWithContext(ctx, fmt.Sprintf("%s.%s", sp, sm)); err == nil {
		p.span = span
	}
	return nil
}

func (p *JaegerPlugin) PostWriteResponse(ctx context.Context, req *protocol.Message, res *protocol.Message, e error) error {
	if p.span != nil {
		p.span.Finish()
	}
	return nil
}
