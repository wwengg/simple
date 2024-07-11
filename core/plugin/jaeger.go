package plugin

import (
	"context"
	"fmt"
	"log"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/smallnest/rpcx/protocol"
)

type JaegerPlugin struct {
	span opentracing.Span
}

func NewJaegerPlugin(t opentracing.Tracer) *JaegerPlugin {
	opentracing.SetGlobalTracer(t)
	return &JaegerPlugin{}
}

func (p *JaegerPlugin) PostReadRequest(ctx context.Context, r *protocol.Message, e error) error {
	sp := r.ServicePath
	sm := r.ServiceMethod

	if sp == "" {
		return nil
	}
	operationName := fmt.Sprintf("rpcx:%s.%s", sp, sm)

	var span opentracing.Span
	var md opentracing.TextMapCarrier
	tracer := opentracing.GlobalTracer()
	if r.Metadata != nil {
		md = r.Metadata
		spanContext, err := tracer.Extract(opentracing.TextMap, md)
		if err != nil && err != opentracing.ErrSpanContextNotFound {
			log.Printf("metadata error %s\n", err)
			return nil
		}
		span = tracer.StartSpan(operationName, ext.RPCServerOption(spanContext))
	} else {
		span = opentracing.StartSpan(operationName)
		md = make(map[string]string)
	}

	err := tracer.Inject(span.Context(), opentracing.TextMap, md)
	if err != nil {
		return nil
	}
	r.Metadata = md
	p.span = span
	return nil
}

func (p *JaegerPlugin) PostWriteResponse(ctx context.Context, req *protocol.Message, res *protocol.Message, e error) error {
	if p.span != nil {
		p.span.Finish()
	}
	return nil
}
