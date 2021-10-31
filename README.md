# api_open_sdk

## 开发须知

必须正确置于gopath下，否则 proto 无法正确生成代码

需要注意env环境变量中需要有gopath/bin

## 安装`proto`

```shell
brew install protobuf
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

## 生成`proto`代码

```shell
export PATH=$PATH:$GOPATH/bin
make
```