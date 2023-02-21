// https://github.com/aliyun-sls/opentelemetry-go-provider-sls/blob/master/provider/sls.go

package aliyunsls

import (
	"context"
	"errors"
	"github.com/sethvargo/go-envconfig"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	otlpTraceGrpc "go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.12.0"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/encoding/gzip"
	"os"
	"strings"
)

const (
	slsProjectHeader         = "x-sls-otel-project"
	slsInstanceIDHeader      = "x-sls-otel-instance-id"
	slsAccessKeyIDHeader     = "x-sls-otel-ak-id"
	slsAccessKeySecretHeader = "x-sls-otel-ak-secret"
	slsSecurityTokenHeader   = "x-sls-otel-token"
)

func parseEnvKeys(c *Config) {
	if c.AttributesEnvKeys == "" {
		return
	}
	envKeys := strings.Split(c.AttributesEnvKeys, "|")
	for _, key := range envKeys {
		key = strings.TrimSpace(key)
		value := os.Getenv(key)
		if value != "" {
			c.resourceAttributes[key] = value
		}
	}
}

// 默认使用本机hostname作为hostname
func getDefaultResource(c *Config) *resource.Resource {
	hostname, _ := os.Hostname()
	return resource.NewWithAttributes(
		semconv.SchemaURL,
		semconv.ServiceNameKey.String(c.ServiceName),
		semconv.HostNameKey.String(hostname),
		semconv.ServiceNamespaceKey.String(c.ServiceNamespace),
		semconv.ServiceVersionKey.String(c.ServiceVersion),
		semconv.ProcessPIDKey.Int(os.Getpid()),
		semconv.ProcessCommandKey.String(os.Args[0]),
	)
}

func mergeResource(c *Config) error {
	var e error
	if c.Resource, e = resource.Merge(getDefaultResource(c), c.Resource); e != nil {
		return e
	}

	r := resource.Environment()
	if c.Resource, e = resource.Merge(c.Resource, r); e != nil {
		return e
	}

	var keyValues []attribute.KeyValue
	for key, value := range c.resourceAttributes {
		keyValues = append(keyValues, attribute.KeyValue{
			Key:   attribute.Key(key),
			Value: attribute.StringValue(value),
		})
	}
	newResource := resource.NewWithAttributes(semconv.SchemaURL, keyValues...)
	if c.Resource, e = resource.Merge(c.Resource, newResource); e != nil {
		return e
	}
	return nil
}

// 初始化Exporter，如果otlpEndpoint传入的值为 stdout，则默认把信息打印到标准输出用于调试
func (c *Config) initOtelExporter(otlpEndpoint string, insecure bool) (trace.SpanExporter, func(), error) {
	var traceExporter trace.SpanExporter
	var err error

	var exporterStop = func() {
		if traceExporter != nil {
			traceExporter.Shutdown(context.Background())
		}
	}

	if otlpEndpoint == "stdout" {
		// 使用Pretty的打印方式
		traceExporter, err = stdouttrace.New(stdouttrace.WithPrettyPrint())
		if err != nil {
			return nil, nil, err
		}
	} else if otlpEndpoint != "" {
		headers := map[string]string{}
		if c.Project != "" && c.InstanceID != "" {
			headers = map[string]string{
				slsProjectHeader:         c.Project,
				slsInstanceIDHeader:      c.InstanceID,
				slsAccessKeyIDHeader:     c.AccessKeyID,
				slsAccessKeySecretHeader: c.AccessKeySecret,
			}
		}

		// 使用GRPC方式导出数据
		traceSecureOption := otlpTraceGrpc.WithTLSCredentials(credentials.NewClientTLSFromCert(nil, ""))
		if insecure {
			traceSecureOption = otlpTraceGrpc.WithInsecure()
		}
		traceExporter, err = otlptrace.New(context.Background(),
			otlpTraceGrpc.NewClient(otlpTraceGrpc.WithEndpoint(otlpEndpoint),
				traceSecureOption,
				otlpTraceGrpc.WithHeaders(headers),
				otlpTraceGrpc.WithCompressor(gzip.Name)))
		if err != nil {
			return nil, nil, err
		}

	}

	return traceExporter, exporterStop, nil
}

// 初始化Traces，默认全量上传
func (c *Config) initTracer(traceExporter trace.SpanExporter, stop func(), config *Config) error {
	if traceExporter == nil {
		return nil
	}
	// 建议使用AlwaysSample全量上传Trace数据，若您的数据太多，可以使用sdktrace.ProbabilitySampler进行采样上传
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(
			traceExporter,
		),
		sdktrace.WithIDGenerator(config.IDGenerator),
		sdktrace.WithResource(c.Resource),
	)
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))
	c.stop = append(c.stop, func() {
		tp.Shutdown(context.Background())
		stop()
	})
	return nil
}

// IsValid check config and return error if config invalid
func (c *Config) IsValid() error {
	if c.ServiceName == "" {
		return errors.New("empty service name")
	}
	if c.ServiceVersion == "" {
		return errors.New("empty service version")
	}
	if (strings.Contains(c.TraceExporterEndpoint, "log.aliyuncs.com") && c.TraceExporterEndpointInsecure) ||
		(strings.Contains(c.MetricExporterEndpoint, "log.aliyuncs.com") && c.MetricExporterEndpointInsecure) {
		return errors.New("insecure grpc is not allowed when send data to sls directly")
	}
	if strings.Contains(c.TraceExporterEndpoint, "log.aliyuncs.com") || strings.Contains(c.MetricExporterEndpoint, "log.aliyuncs.com") {
		if c.Project == "" || c.InstanceID == "" || c.AccessKeyID == "" || c.AccessKeySecret == "" {
			return errors.New("empty project, instanceID, accessKeyID or accessKeySecret when send data to sls directly")
		}
		if strings.ContainsAny(c.Project, "${}") ||
			strings.ContainsAny(c.InstanceID, "${}") ||
			strings.ContainsAny(c.AccessKeyID, "${}") ||
			strings.ContainsAny(c.AccessKeySecret, "${}") {
			return errors.New("invalid project, instanceID, accessKeyID or accessKeySecret when send data to sls directly, you should replace these parameters with actual values")
		}
	}
	return nil
}

// NewConfig create a config
func NewConfig(opts ...Option) (*Config, error) {
	var c Config

	// 1. load env config
	envError := envconfig.Process(context.Background(), &c)
	if envError != nil {
		return nil, envError
	}

	// 2. load code config
	for _, opt := range opts {
		opt(&c)
	}

	// 3. merge resource
	parseEnvKeys(&c)
	mergeResource(&c)
	return &c, c.IsValid()
}

// Start 初始化OpenTelemetry SDK，需要把 ${endpoint} 替换为实际的地址
// 如果填写为stdout则为调试默认，数据将打印到标准输出
func Start(c *Config) error {
	if c.errorHandler != nil {
		otel.SetErrorHandler(c.errorHandler)
	}
	traceExporter, traceExpStop, err := c.initOtelExporter(c.TraceExporterEndpoint, c.TraceExporterEndpointInsecure)
	if err != nil {
		return err
	}
	err = c.initTracer(traceExporter, traceExpStop, c)
	return err
}

// Shutdown 优雅关闭，将OpenTelemetry SDK内存中的数据发送到服务端
func Shutdown(c *Config) {
	for _, stop := range c.stop {
		stop()
	}
}
