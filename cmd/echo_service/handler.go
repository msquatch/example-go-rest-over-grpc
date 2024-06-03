package main

import (
	// Built-in/core modules.
	"context"

	// Third-party modules.

	// Generated code.
	echo "my/echo"
	echo_service "my/service"
)

// Structure used for the handler object.
type EchoServer struct {
	echo_service.UnimplementedEchoServer
}

// Implements the EchoString service method defined in service-echo.proto.
func (s *EchoServer) EchoString(
	ctx context.Context,
	in *echo.StringRequest,
) (*echo.StringReply, error) {
	return &echo.StringReply{Value: in.Value}, nil
}

// Implements the EchoInt service method defined in service-echo.proto.
func (s *EchoServer) EchoInt(
	ctx context.Context,
	in *echo.IntRequest,
) (*echo.IntReply, error) {
	return &echo.IntReply{Value: in.Value}, nil
}
