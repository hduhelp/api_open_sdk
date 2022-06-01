# api_open_sdk

![Goproxy.cn](https://goproxy.cn/stats/github.com/hduhelp/api_open_sdk/badges/download-count.svg)
![buf lint](https://github.com/hduhelp/api_open_sdk/actions/workflows/buf.yml/badge.svg)

## 前置步骤

```shell
//安装 编译所需的插件
//包括 buf protoc-gen-grpc-gateway protoc-gen-openapiv2 protoc-gen-go-grpc protoc-gen-go
//在 tools.go 中控制 要安装的插件列表，使用 go module 管理插件版本
make install-tools
```

## 开发流程

1. 修改 proto 文件
2. 运行 `buf generate`
3. 开启 swagger-ui 实时预览

## 预览 swagger 效果 (多数情况下没必要预览) (感觉还有更简单的)

### 公开文档

在 https://github.com/hduhelp/api/config.json 中添加对应 OpenAPI 文件，提交后会自动拉取api_open_sdk最新文档

可在 https://hduhelp.github.io/api/ 页面查看接口文档

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

- 有效：rpc 方法前，Response message 前，message内部字段前
- 无效（不会显示到 swagger 中）：Request message 前
