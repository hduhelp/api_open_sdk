module github.com/hduhelp/api_open_sdk

go 1.22.0

toolchain go1.22.7

require (
	github.com/avct/uasurfer v0.0.0-20191028135549-26b5daa857f1
	github.com/bufbuild/buf v1.45.0
	github.com/eko/gocache/v2 v2.3.1
	github.com/felixge/httpsnoop v1.0.4
	github.com/gin-gonic/gin v1.9.1
	github.com/go-redis/redis/v8 v8.11.5
	github.com/grpc-ecosystem/go-grpc-middleware v1.4.0
	github.com/grpc-ecosystem/go-grpc-prometheus v1.2.0
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.22.0
	github.com/parnurzeal/gorequest v0.2.16
	github.com/patrickmn/go-cache v2.1.0+incompatible
	github.com/pkg/errors v0.9.1
	github.com/rabbitmq/amqp091-go v1.9.0
	github.com/samber/lo v1.38.1
	github.com/satori/go.uuid v1.2.0
	github.com/sethvargo/go-envconfig v0.9.0
	github.com/silenceper/wechat/v2 v2.1.6
	github.com/soheilhy/cmux v0.1.5
	go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.46.1
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.55.0
	go.opentelemetry.io/otel v1.30.0
	go.opentelemetry.io/otel/exporters/otlp/otlptrace v1.19.0
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc v1.19.0
	go.opentelemetry.io/otel/exporters/stdout/stdouttrace v1.15.1
	go.opentelemetry.io/otel/sdk v1.28.0
	go.opentelemetry.io/otel/trace v1.30.0
	go.uber.org/zap v1.27.0
	google.golang.org/genproto/googleapis/api v0.0.0-20240930140551-af27646dc61f
	google.golang.org/grpc v1.67.1
	google.golang.org/grpc/cmd/protoc-gen-go-grpc v1.5.1
	google.golang.org/protobuf v1.35.1
)

require (
	buf.build/gen/go/bufbuild/bufplugin/protocolbuffers/go v1.34.2-20240928190436-5e8abcfd7a7e.2 // indirect
	buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go v1.34.2-20240920164238-5a7b106cbb87.2 // indirect
	buf.build/gen/go/bufbuild/registry/connectrpc/go v1.17.0-20240925012807-1610ffa05635.1 // indirect
	buf.build/gen/go/bufbuild/registry/protocolbuffers/go v1.34.2-20240925012807-1610ffa05635.2 // indirect
	buf.build/gen/go/pluginrpc/pluginrpc/protocolbuffers/go v1.34.2-20240828222655-5345c0a56177.2 // indirect
	buf.build/go/bufplugin v0.5.0 // indirect
	buf.build/go/protoyaml v0.2.0 // indirect
	buf.build/go/spdx v0.2.0 // indirect
	cloud.google.com/go/compute/metadata v0.5.0 // indirect
	connectrpc.com/connect v1.17.0 // indirect
	connectrpc.com/otelconnect v0.7.1 // indirect
	github.com/Azure/go-ansiterm v0.0.0-20230124172434-306776ec8161 // indirect
	github.com/Microsoft/go-winio v0.6.2 // indirect
	github.com/Microsoft/hcsshim v0.12.7 // indirect
	github.com/XiaoMi/pegasus-go-client v0.0.0-20210427083443-f3b6b08bc4c2 // indirect
	github.com/antlr4-go/antlr/v4 v4.13.1 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/bradfitz/gomemcache v0.0.0-20220106215444-fb4bf637b56d // indirect
	github.com/bufbuild/protocompile v0.14.1 // indirect
	github.com/bufbuild/protoplugin v0.0.0-20240911180120-7bb73e41a54a // indirect
	github.com/bufbuild/protovalidate-go v0.7.2 // indirect
	github.com/bytedance/sonic v1.9.1 // indirect
	github.com/cenkalti/backoff/v4 v4.3.0 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/chenzhuoyu/base64x v0.0.0-20221115062448-fe3a3abad311 // indirect
	github.com/containerd/cgroups/v3 v3.0.3 // indirect
	github.com/containerd/containerd v1.7.22 // indirect
	github.com/containerd/continuity v0.4.3 // indirect
	github.com/containerd/errdefs v0.2.0 // indirect
	github.com/containerd/log v0.1.0 // indirect
	github.com/containerd/platforms v0.2.1 // indirect
	github.com/containerd/stargz-snapshotter/estargz v0.15.1 // indirect
	github.com/containerd/ttrpc v1.2.5 // indirect
	github.com/containerd/typeurl/v2 v2.2.0 // indirect
	github.com/cpuguy83/go-md2man/v2 v2.0.5 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/distribution/reference v0.6.0 // indirect
	github.com/docker/cli v27.3.1+incompatible // indirect
	github.com/docker/distribution v2.8.3+incompatible // indirect
	github.com/docker/docker v27.3.1+incompatible // indirect
	github.com/docker/docker-credential-helpers v0.8.2 // indirect
	github.com/docker/go-connections v0.5.0 // indirect
	github.com/docker/go-units v0.5.0 // indirect
	github.com/felixge/fgprof v0.9.5 // indirect
	github.com/gabriel-vasile/mimetype v1.4.2 // indirect
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/go-chi/chi/v5 v5.1.0 // indirect
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/go-playground/validator/v10 v10.14.0 // indirect
	github.com/go-task/slim-sprig/v3 v3.0.0 // indirect
	github.com/goccy/go-json v0.10.2 // indirect
	github.com/gofrs/flock v0.12.1 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/google/cel-go v0.21.0 // indirect
	github.com/google/go-containerregistry v0.20.2 // indirect
	github.com/google/pprof v0.0.0-20241001023024-f4c0cfd0cf1d // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/jdx/go-netrc v1.0.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/compress v1.17.10 // indirect
	github.com/klauspost/cpuid/v2 v2.2.4 // indirect
	github.com/klauspost/pgzip v1.2.6 // indirect
	github.com/leodido/go-urn v1.2.4 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	github.com/moby/docker-image-spec v1.3.1 // indirect
	github.com/moby/locker v1.0.1 // indirect
	github.com/moby/patternmatcher v0.6.0 // indirect
	github.com/moby/sys/mount v0.3.4 // indirect
	github.com/moby/sys/mountinfo v0.7.2 // indirect
	github.com/moby/sys/sequential v0.6.0 // indirect
	github.com/moby/sys/user v0.3.0 // indirect
	github.com/moby/sys/userns v0.1.0 // indirect
	github.com/moby/term v0.5.0 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/morikuni/aec v1.0.0 // indirect
	github.com/munnerz/goautoneg v0.0.0-20191010083416-a7dc8b61c822 // indirect
	github.com/onsi/ginkgo/v2 v2.20.2 // indirect
	github.com/opencontainers/go-digest v1.0.0 // indirect
	github.com/opencontainers/image-spec v1.1.0 // indirect
	github.com/opencontainers/runtime-spec v1.2.0 // indirect
	github.com/pegasus-kv/thrift v0.13.0 // indirect
	github.com/pelletier/go-toml/v2 v2.0.8 // indirect
	github.com/pkg/browser v0.0.0-20240102092130-5ac0b6a4141c // indirect
	github.com/pkg/profile v1.7.0 // indirect
	github.com/prometheus/client_golang v1.20.2 // indirect
	github.com/prometheus/client_model v0.6.1 // indirect
	github.com/prometheus/common v0.55.0 // indirect
	github.com/prometheus/procfs v0.15.1 // indirect
	github.com/quic-go/qpack v0.5.1 // indirect
	github.com/quic-go/quic-go v0.47.0 // indirect
	github.com/rs/cors v1.11.1 // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	github.com/segmentio/asm v1.2.0 // indirect
	github.com/segmentio/encoding v0.4.0 // indirect
	github.com/sirupsen/logrus v1.9.3 // indirect
	github.com/spf13/cast v1.5.0 // indirect
	github.com/spf13/cobra v1.8.1 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/stoewer/go-strcase v1.3.0 // indirect
	github.com/tetratelabs/wazero v1.8.1 // indirect
	github.com/twitchyliquid64/golang-asm v0.15.1 // indirect
	github.com/ugorji/go/codec v1.2.11 // indirect
	github.com/vbatts/tar-split v0.11.6 // indirect
	github.com/vmihailenco/msgpack v4.0.4+incompatible // indirect
	go.lsp.dev/jsonrpc2 v0.10.0 // indirect
	go.lsp.dev/pkg v0.0.0-20210717090340-384b27a52fb2 // indirect
	go.lsp.dev/protocol v0.12.0 // indirect
	go.lsp.dev/uri v0.3.0 // indirect
	go.opencensus.io v0.24.0 // indirect
	go.opentelemetry.io/otel/metric v1.30.0 // indirect
	go.opentelemetry.io/proto/otlp v1.0.0 // indirect
	go.uber.org/atomic v1.11.0 // indirect
	go.uber.org/mock v0.4.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	go.uber.org/zap/exp v0.1.1-0.20240913022758-ede8e1888f83 // indirect
	golang.org/x/arch v0.3.0 // indirect
	golang.org/x/crypto v0.28.0 // indirect
	golang.org/x/exp v0.0.0-20241004190924-225e2abe05e6 // indirect
	golang.org/x/mod v0.21.0 // indirect
	golang.org/x/net v0.30.0 // indirect
	golang.org/x/sync v0.8.0 // indirect
	golang.org/x/sys v0.26.0 // indirect
	golang.org/x/term v0.25.0 // indirect
	golang.org/x/text v0.19.0 // indirect
	golang.org/x/tools v0.26.0 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240930140551-af27646dc61f // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.0.0 // indirect
	gopkg.in/tomb.v2 v2.0.0-20161208151619-d5d1b5820637 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	k8s.io/apimachinery v0.26.2 // indirect
	moul.io/http2curl v1.0.0 // indirect
	pluginrpc.com/pluginrpc v0.5.0 // indirect
)
