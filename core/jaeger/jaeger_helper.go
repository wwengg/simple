package jaeger

import (
	"io"
	"time"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
)

func NewTracer(servicename string, addr string) (opentracing.Tracer, io.Closer, error) {
	cfg := jaegercfg.Configuration{
		ServiceName: servicename,
		/*
		   jaeger.SamplerTypeConst:          全量采集，采样率设置0,1 分别对应打开和关闭
		   jaeger.SamplerTypeProbabilistic:  概率采集，默认万份之一，0~1之间取值，
		   jaeger.SamplerTypeRateLimiting:   限速采集，每秒只能采集一定量的数据
		   jaeger.SamplerTypeRemote:         一种动态采集策略，根据当前系统的访问量调节采集策略
		*/
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans:            true,
			BufferFlushInterval: 1 * time.Second,
		},
	}

	sender, err := jaeger.NewUDPTransport(addr, 0)
	if err != nil {
		return nil, nil, err
	}

	reporter := jaeger.NewRemoteReporter(sender)
	// Initialize tracer with a logger and a metrics factory
	tracer, closer, err := cfg.NewTracer(
		jaegercfg.Reporter(reporter),
	)

	return tracer, closer, err
}
