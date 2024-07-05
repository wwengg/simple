package plugin

import (
	"context"
	"fmt"
	"log"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/smallnest/rpcx/protocol"
	"github.com/smallnest/rpcx/share"
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
	if span, _, err := generateSpanWithContext(ctx, fmt.Sprintf("%s.%s", sp, sm)); err == nil {
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

// 只适用于 jaeger
// 原理用 context 传递 "__req_metadata":""uber-trace-id -> 6f8b8a1101b0124f:6f8b8a1101b0124f:0000000000000000:1
//
//	uber-trace-id traceID : spanID : parentID: sampled bool
func generateSpanWithContext(ctx context.Context, operationName string) (opentracing.Span, context.Context, error) {
	md := ctx.Value(share.ReqMetaDataKey) // share.ReqMetaDataKey 固定值 "__req_metadata"  可自定义
	var span opentracing.Span

	tracer := opentracing.GlobalTracer()

	if md != nil {
		carrier := opentracing.TextMapCarrier(md.(map[string]string))
		spanContext, err := tracer.Extract(opentracing.TextMap, carrier)
		if err != nil && err != opentracing.ErrSpanContextNotFound {
			log.Printf("metadata error %s\n", err)
			return nil, nil, err
		}
		span = tracer.StartSpan(operationName, ext.RPCServerOption(spanContext))
	} else {
		span = opentracing.StartSpan(operationName)
	}

	metadata := opentracing.TextMapCarrier(make(map[string]string))
	err := tracer.Inject(span.Context(), opentracing.TextMap, metadata)
	if err != nil {
		return nil, nil, err
	}
	//把metdata 携带的 traceid,spanid,parentSpanid 放入 context
	ctx = context.WithValue(context.Background(), share.ReqMetaDataKey, (map[string]string)(metadata))
	return span, ctx, nil
}
