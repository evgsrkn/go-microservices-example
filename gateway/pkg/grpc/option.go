package rpc

import (
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type Option = func() *grpc.ServerOption

func WithZapLogger(logger *zap.Logger) Option {
	return func() *grpc.ServerOption {
		o := []grpc_zap.Option{
			grpc_zap.WithLevels(grpc_zap.DefaultCodeToLevel),
		}
		// Make sure that log statements internal to gRPC library are logged using the zapLogger as well.
		grpc_zap.ReplaceGrpcLogger(logger)

		opt := grpc_middleware.WithUnaryServerChain(
			grpc_ctxtags.UnaryServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
			grpc_zap.UnaryServerInterceptor(logger, o...),
		)

		return &opt
	}
}
