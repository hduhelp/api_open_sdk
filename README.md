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
2. 运行 `buf generate`
3. 开启 swagger-ui 实时预览

## 预览 swagger 效果 (感觉还有更简单的)
1. 安装 swagger-ui
```shell
# 新建目录
git clone https://github.com/swagger-api/swagger-ui.git
cd swagger-ui
npm install
npm run dev
```
2. 安装 http-server
```shell
npm install --global http-server
# 切回 api_open_sdk 根目录
cd docs/swagger-ui
http-server --cors
```
3. 访问 swagger-ui
访问 http://localhost:3200/  
在地址栏内输入 json 的相对路径，如：
http://localhost:8080/campusapis/staff/v1/freshman.swagger.json
点击 Explore 按钮，即可看到 swagger 文档

4. 每次改动 proto 后需要重新运行 `buf generate`，并且「清空缓存并进行硬刷新」网页

## 接口注释

有效：rpc 方法前，Response message 前，message内部字段前
无效（不会显示到 swagger 中）：Request message 前

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