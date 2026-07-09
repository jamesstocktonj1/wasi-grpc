package main

import (
	"context"
	"fmt"
	"wasi-grpc/gen"
)

type server struct {
	gen.UnimplementedGreeterServer
}

var _ gen.GreeterServer = (*server)(nil)

func (s *server) Greet(ctx context.Context, req *gen.Request) (*gen.Response, error) {
	return &gen.Response{
		Message: fmt.Sprintf("Hello, %s!", req.Name),
	}, nil
}
