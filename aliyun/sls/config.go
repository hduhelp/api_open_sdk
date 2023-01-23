// https://github.com/aliyun-sls/opentelemetry-go-provider-sls/blob/master/provider/sls.go

package aliyunsls

import (
	"fmt"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"time"
)

// Option configures the sls otel provider
type Option func(*Config)

// WithMetricExporterEndpoint configures the endpoint for sending metrics via OTLP
// 配置Metric的输出地址，如果配置为空则禁用Metric功能，配置为stdout则打印到标准输出用于测试
func WithMetricExporterEndpoint(url string) Option {
	return func(c *Config) {
		c.MetricExporterEndpoint = url
	}
}

// WithTraceExporterEndpoint configures the endpoint for sending traces via OTLP
// 配置Trace的输出地址，如果配置为空则禁用Trace功能，配置为stdout则打印到标准输出用于测试
func WithTraceExporterEndpoint(url string) Option {
	return func(c *Config) {
		c.TraceExporterEndpoint = url
	}
}

// WithServiceName configures a "service.name" resource label
// 配置服务名称
func WithServiceName(name string) Option {
	return func(c *Config) {
		c.ServiceName = name
	}
}

func WithServiceNamespace(namespace string) Option {
	return func(c *Config) {
		c.ServiceNamespace = namespace
	}
}

// WithServiceVersion configures a "service.version" resource label
// 配置版本号
func WithServiceVersion(version string) Option {
	return func(c *Config) {
		c.ServiceVersion = version
	}
}

// WithTraceExporterInsecure permits connecting to the trace endpoint without a certificate
// 配置是否禁用SSL，如果输出到SLS，则必须打开SLS
func WithTraceExporterInsecure(insecure bool) Option {
	return func(c *Config) {
		c.TraceExporterEndpointInsecure = insecure
	}
}

// WithMetricExporterInsecure permits connecting to the metric endpoint without a certificate
// 配置是否禁用SSL，如果输出到SLS，则必须打开SLS
func WithMetricExporterInsecure(insecure bool) Option {
	return func(c *Config) {
		c.MetricExporterEndpointInsecure = insecure
	}
}

// WithResourceAttributes configures attributes on the resource
// 配置上传附加的一些tag信息，例如环境、可用区等
func WithResourceAttributes(attributes map[string]string) Option {
	return func(c *Config) {
		c.resourceAttributes = attributes
	}
}

// WithResource configures attributes on the resource
// 配置上传附加的一些tag信息，例如环境、可用区等
func WithResource(resource *resource.Resource) Option {
	return func(c *Config) {
		c.Resource = resource
	}
}

// WithErrorHandler Configures a global error handler to be used throughout an OpenTelemetry instrumented project.
// See "go.opentelemetry.io/otel"
// 配置OpenTelemetry错误处理函数
func WithErrorHandler(handler otel.ErrorHandler) Option {
	return func(c *Config) {
		c.errorHandler = handler
	}
}

// WithMetricReportingPeriod configures the metric reporting period,
// how often the controller collects and exports metric data.
// 配置Metric导出间隔，默认为30s
func WithMetricReportingPeriod(p time.Duration) Option {
	return func(c *Config) {
		c.MetricReportingPeriod = fmt.Sprint(p)
	}
}

// WithSLSConfig configures sls project, instanceID, accessKeyID, accessKeySecret to send data to sls directly
// 配置输出到SLS的信息，包括 project, instanceID, accessKeyID, accessKeySecret
func WithSLSConfig(project, instanceID, accessKeyID, accessKeySecret string) Option {
	return func(c *Config) {
		c.Project, c.InstanceID, c.AccessKeyID, c.AccessKeySecret = project, instanceID, accessKeyID, accessKeySecret
	}
}

func WithIDGenerator(generator sdktrace.IDGenerator) Option {
	return func(config *Config) {
		if generator != nil {
			config.IDGenerator = generator
		}
	}
}

// Config configure for sls otel
type Config struct {
	TraceExporterEndpoint          string `env:"SLS_OTEL_TRACE_ENDPOINT,default=stdout"`
	TraceExporterEndpointInsecure  bool   `env:"SLS_OTEL_TRACE_INSECURE,default=false"`
	MetricExporterEndpoint         string `env:"SLS_OTEL_METRIC_ENDPOINT,default=stdout"`
	MetricExporterEndpointInsecure bool   `env:"SLS_OTEL_METRIC_INSECURE,default=false"`
	MetricReportingPeriod          string `env:"SLS_OTEL_METRIC_EXPORT_PERIOD,default=30s"`
	ServiceName                    string `env:"SLS_OTEL_SERVICE_NAME"`
	ServiceNamespace               string `env:"SLS_OTEL_SERVICE_NAMESPACE"`
	ServiceVersion                 string `env:"SLS_OTEL_SERVICE_VERSION,default=v0.1.0"`
	Project                        string `env:"SLS_OTEL_PROJECT"`
	InstanceID                     string `env:"SLS_OTEL_INSTANCE_ID"`
	AccessKeyID                    string `env:"SLS_OTEL_ACCESS_KEY_ID"`
	AccessKeySecret                string `env:"SLS_OTEL_ACCESS_KEY_SECRET"`
	AttributesEnvKeys              string `env:"SLS_OTEL_ATTRIBUTES_ENV_KEYS"`
	IDGenerator                    sdktrace.IDGenerator

	Resource *resource.Resource

	resourceAttributes map[string]string
	errorHandler       otel.ErrorHandler
	stop               []func()
}
