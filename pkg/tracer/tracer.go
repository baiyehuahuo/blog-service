package tracer

import (
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go/config"
	"io"
	"time"
)

func NewJaegerTracer(serviceName, agentHostPort string) (opentracing.Tracer, io.Closer, error) {
	// jaeger client 的配置项 设置应用的基本信息
	cfg := &config.Configuration{
		ServiceName: serviceName,
		// 固定采样、对所有数据进行采样
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		// 是否开启 LoggingReporter 刷新缓冲区频率 上报地址
		Reporter: &config.ReporterConfig{
			LogSpans:            true,
			BufferFlushInterval: time.Second,
			LocalAgentHostPort:  agentHostPort,
		},
	}
	tracer, closer, err := cfg.NewTracer()
	if err != nil {
		return nil, nil, err
	}
	// 设置全局 Tracer
	opentracing.SetGlobalTracer(tracer)
	return tracer, closer, nil
}
