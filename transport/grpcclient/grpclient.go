package grpcclient

import (
	"context"
	"fmt"

	"github.com/underbek/examples-go/logger"
	trace "go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"

	grpcZap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpcPrometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
)

func NewConnection(logger *logger.Logger, address string) (*grpc.ClientConn, error) {
	opts := []grpc.DialOption{
		grpc.WithStreamInterceptor(trace.StreamClientInterceptor()),
		grpc.WithUnaryInterceptor(trace.UnaryClientInterceptor()),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithChainUnaryInterceptor(
			grpcPrometheus.UnaryClientInterceptor,
			grpcZap.UnaryClientInterceptor(
				logger.Named("grpc-client").Internal().(*zap.Logger),
				grpcZap.WithLevels(func(code codes.Code) zapcore.Level {
					return zapcore.DebugLevel
				}),
			),
			CustomPayloadUnaryClientInterceptor(
				logger.Named("grpc-client-payload").Internal().(*zap.Logger),
				func(ctx context.Context, fullMethodName string) bool {
					return logger.Internal().(*zap.Logger).Core().Enabled(zapcore.DebugLevel)
				},
			),
		),
	}

	conn, err := grpc.Dial(
		address,
		opts...,
	)

	if err != nil {
		return nil, fmt.Errorf("create connnection: %w", err)
	}

	return conn, nil
}
