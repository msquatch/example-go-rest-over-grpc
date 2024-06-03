# example-go-rest-over-grpc

Example of providing a REST API on top of gRPC in Go.

## gRPC

Example calls:

```sh
grpcurl -plaintext -import-path . -proto service-echo.proto -proto echo.proto -d '{"value": "foo"}' 127.0.0.1:9000 Echo/EchoString
```

```sh
grpcurl -plaintext -import-path . -proto service-echo.proto -proto echo.proto -d '{"value": 42}' 127.0.0.1:9000 Echo/EchoInt
```

## REST

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
