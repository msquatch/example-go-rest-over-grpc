package main

import (
	// Built-in/core modules.

	"context"
	"fmt"
	slog "log/slog"
	"net"
	"net/http"
	"os"
	"strings"

	// Third-party modules.
	gwruntime "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	cmux "github.com/soheilhy/cmux"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/encoding/protojson"

	// Generated code.
	echo_service "my/service"
)

type rest_handler_func func(ctx context.Context, mux *gwruntime.ServeMux, endpoint string, opts []grpc.DialOption) (err error)

func configure_rest_and_serve(
	conf *Config,
	listener net.Listener,
	grpc_server *grpc.Server,
	reg_rest_handler rest_handler_func,
) error {
	gwmux := gwruntime.NewServeMux(
		gwruntime.WithMarshalerOption(gwruntime.MIMEWildcard, &gwruntime.JSONPb{
			MarshalOptions: protojson.MarshalOptions{
				UseProtoNames: true,
			},
		}),
		// Make sure all the headers are passed along.
		gwruntime.WithIncomingHeaderMatcher(func(key string) (string, bool) {
			return strings.ToLower(key), true
		}),
	)

	err := reg_rest_handler(context.Background(), gwmux, conf.Listen,
		[]grpc.DialOption{
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		})
	if err != nil {
		return fmt.Errorf("couldn't register REST handler: %w", err)
	}

	gw_server := &http.Server{
		Addr:    conf.Listen,
		Handler: gwmux,
	}

	m := cmux.New(listener)
	// Match HTTP 2.x requests.
	http2_listener := m.Match(cmux.HTTP2())
	// Match HTTP 1.x requests.
	http1_listener := m.Match(cmux.HTTP1Fast())

	go gw_server.Serve(http1_listener)
	go grpc_server.Serve(http2_listener)

	slog.Info("Starting server", "listen", conf.Listen)
	return m.Serve()
}

func serve(conf *Config) error {
	grpc_server := grpc.NewServer()
	echo_service.RegisterEchoServer(grpc_server, &EchoServer{})

	listener, err := net.Listen("tcp", conf.Listen)
	if err != nil {
		return fmt.Errorf("server couldn't listen on %q: %w",
			conf.Listen, err)
	}

	// err = grpc_server.Serve(listener)
	err = configure_rest_and_serve(conf, listener, grpc_server,
		echo_service.RegisterEchoHandlerFromEndpoint)

	if err != nil {
		slog.Error("couldn't start server", "err", err)
		os.Exit(-1)
	}

	return nil
}

func main() {
	log_level := new(slog.LevelVar)
	log_level.Set(slog.LevelInfo)
	slog.SetDefault(slog.New(slog.NewTextHandler(
		os.Stderr, &slog.HandlerOptions{Level: log_level}),
	))

	conf, err := parse_cmd_opts()
	if err != nil {
		slog.Error("error with command-line parameters", "err", err)
		os.Exit(-1)
	}

	err = serve(conf)
	if err != nil {
		slog.Error("error starting service", "err", err)
		os.Exit(-1)
	}
}
