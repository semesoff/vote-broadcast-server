package middleware

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"log"
	"time"
)

func UnaryLoggingInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	startTime := time.Now()

	response, err := handler(ctx, req)

	duration := time.Since(startTime).Round(time.Millisecond)
	st, _ := status.FromError(err)
	log.Printf("Request - Method: %s, Duration: %s, Error: %v", info.FullMethod, duration, st)

	return response, err
}
