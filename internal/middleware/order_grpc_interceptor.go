package middleware

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"time"
)

// general unary interceptor function to handle auth per RPC call as well as logging
func unaryInterceptor(ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (interface{}, error) {
	start := time.Now()

	h, err := handler(ctx, req)

	//logging
	log.Printf("request - Method:%s\tDuration:%s\tError:%v\n",
		info.FullMethod,
		time.Since(start),
		err)

	return h, err
}