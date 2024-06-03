# example-go-rest-over-grpc

Example of providing a REST API on top of gRPC in Go. The configuration and code provided implement a REST API on top of a simple gRPC service. Both services listen on the same port. Under the hood,when a REST request is received, it is converted to a gRPC request and sent over a socket.

## Examples

### gRPC

Example calls:

```sh
grpcurl -plaintext -import-path . -proto service-echo.proto -proto echo.proto -d '{"value": "foo"}' 127.0.0.1:9000 Echo/EchoString
```

```sh
grpcurl -plaintext -import-path . -proto service-echo.proto -proto echo.proto -d '{"value": 42}' 127.0.0.1:9000 Echo/EchoInt
```

### REST

Various styles of REST styles are supported by the specification in `echo.yaml`:

```sh
# Extract all fields from the JSON-encoded payload. The URL path is the same
# as for gRPC.
curl -d '{"value": "foo"}' "http://127.0.0.1:9000/Echo/EchoString"
```

```sh
# Extract all fields from the JSON-encoded payload. The path is custom.
curl -d '{"value": "foo"}' "http://127.0.0.1:9000/echo/string"
```

```sh
# Extract data from the POST body as URL-encoded fields.
curl -d 'value=foo' "http://127.0.0.1:9000/echo/stringquery"
```

```sh
# Extract all fields from the query string.
curl "http://127.0.0.1:9000/echo/string?value=foo"
```

```sh
# Extract the `value` field from the URL path.
curl "http://127.0.0.1:9000/echo/string/foo"
```

## Dependencies

In addition to the `make` utility, the following dependencies are required to build a working gRPC service in Go with a REST gateway on top.

* [buf](https://buf.build/docs/installation): protobuf build and management tool.
* [protoc](https://grpc.io/docs/protoc-installation/): the protobuf compiler.
* Go protobuf compiler plugins:
  * [protoc-gen-go](https://grpc.io/docs/languages/go/quickstart/): protobuf compiler/code generator for Go.
  * [protoc-gen-go-grpc](https://grpc.io/docs/languages/go/quickstart/): protobuf compiler/code generator for gRPC services in Go.
* [protoc-gen-grpc-gateway](https://github.com/grpc-ecosystem/grpc-gateway?tab=readme-ov-file): gRPC Gateway plugin for the protobuf compiler (generates Go code).
