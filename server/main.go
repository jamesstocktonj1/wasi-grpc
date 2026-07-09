package main

import (
	"github.com/jamesstocktonj1/componentize-sdk/p3/cli"
	"github.com/jamesstocktonj1/componentize-sdk/p3/net/socket"
)

func init() {
	cli.SetRunE(run)
}

func run() error {
	listen, err := socket.Listen("tcp", "127.0.0.1:8040")
	if err != nil {
		return err
	}

	for {
		conn, err := listen.Accept()
		if err != nil {
			return err
		}
		defer conn.Close()
	}
}

func main() {}
