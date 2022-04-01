# api_open_sdk

## 前置步骤

```shell
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway 
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 
go install google.golang.org/protobuf/cmd/protoc-gen-go 
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
```
### For Windows

1. 安装 [scoop](https://scoop.sh/) : ``iwr -useb get.scoop.sh | iex`` (注意不要使用管理员身份运行)
2. 安装 [buf](https://docs.buf.build/installation#scoop) :``scoop install buf``

## 开发流程

1. 修改 proto 文件
2. 运行 ``buf generate``

## 开发须知 old

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