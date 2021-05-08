# protoc-gen-go-gin

从 protobuf 文件中生成使用 gin 的 http rpc 服务，同时兼容 gRPC 服务。

- 支持rest参数
- 支持url query参数
- 支持post/put http body json


请确保安装了以下依赖:

- [go 1.16.4](https://golang.org/dl/)
- [protoc](https://github.com/protocolbuffers/protobuf)
- [protoc-gen-go](https://github.com/protocolbuffers/protobuf-go)

安装方法

```bash
go install github.com/fzpanxi/protoc-gen-go-gin@latest
```

## 使用说明

Demo: [examples](./examples)

### proto 文件约定

 google.api.http option 指定路由，可以通过添加 additional_bindings 使一个 rpc 方法对应多个路由

```protobuf

service HelloService {
  rpc GetHello (GetHelloRequest) returns (GetHelloReply)  {
    option (google.api.http) = {
      get: "/v1/hello/{name}"
      additional_bindings {
        get: "/v1/hello_word/{name}"
      }
    };
  }
}

message GetHelloRequest {
  string name = 1;
}


message GetHelloReply {
  string message = 1;
}

```
### 文件生成

```bash
windows:
  protoc -I=%GOPATH%/src -I ./examples  --go-grpc_out=./examples --go-grpc_opt=paths=source_relative examples/hello/api/hello/v1/hello.proto
  protoc -I=%GOPATH%/src -I ./examples  --go_out=./examples --go_opt=paths=source_relative examples/hello/api/hello/v1/hello.proto
  protoc -I=%GOPATH%/src -I ./examples  --go-gin_out=./examples --go-gin_opt=paths=source_relative examples/hello/api/hello/v1/hello.proto
```
```bash
linux/mac/unix:
  protoc -I=$GOPATH/src -I ./examples  --go-grpc_out=./examples --go-grpc_opt=paths=source_relative examples/hello/api/hello/v1/hello.proto
  protoc -I=$GOPATH/src -I ./examples  --go_out=./examples --go_opt=paths=source_relative examples/hello/api/hello/v1/hello.proto
  protoc -I=$GOPATH/src -I ./examples  --go-gin_out=./examples --go-gin_opt=paths=source_relative examples/hello/api/hello/v1/hello.proto
```