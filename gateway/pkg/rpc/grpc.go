package rpc

import "google.golang.org/grpc"

func New(opts ...Option) *grpc.Server {
	var options []grpc.ServerOption
	for _, option := range opts {
		options = append(options, *option())
	}

	return grpc.NewServer(options...)
}
