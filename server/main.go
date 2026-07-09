package main

import (
	"wasi-grpc/gen"

	"github.com/jamesstocktonj1/componentize-sdk/p3/cli"
	"github.com/jamesstocktonj1/componentize-sdk/p3/net/socket"
	"google.golang.org/grpc"
)

func init() {
	cli.SetRunE(run)
}

func run() error {
	listen, err := socket.Listen("tcp", "127.0.0.1:8040")
	if err != nil {
		return err
	}

	sgrpc := grpc.NewServer(nil)
	gen.RegisterGreeterServer(sgrpc, &server{})
	return sgrpc.Serve(listen)
}

func main() {}
