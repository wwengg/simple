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
	operationName := fmt.Sprintf("rpcx:%s.%s", sp, sm)
	md := ctx.Value(share.ReqMetaDataKey) // share.ReqMetaDataKey 固定值 "__req_metadata"  可自定义
	var span opentracing.Span

	tracer := opentracing.GlobalTracer()
	if md == nil {
		md = r.Metadata
	}
	if md != nil {
		carrier := opentracing.TextMapCarrier(md.(map[string]string))
		spanContext, err := tracer.Extract(opentracing.TextMap, carrier)
		if err != nil && err != opentracing.ErrSpanContextNotFound {
			log.Printf("metadata error %s\n", err)
			return nil
		}
		span = tracer.StartSpan(operationName, ext.RPCServerOption(spanContext))
	} else {
		span = opentracing.StartSpan(operationName)
	}

	metadata := opentracing.TextMapCarrier(make(map[string]string))
	err := tracer.Inject(span.Context(), opentracing.TextMap, metadata)
	if err != nil {
		return nil
	}

	ctx = context.WithValue(ctx, share.ReqMetaDataKey, (map[string]string)(metadata))
	p.span = span
	return nil
}

func (p *JaegerPlugin) PostWriteResponse(ctx context.Context, req *protocol.Message, res *protocol.Message, e error) error {
	if p.span != nil {
		p.span.Finish()
	}
	return nil
}
