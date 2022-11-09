package grpcmiddleware

import (
	"context"
	"log"
	"path"
	"time"

	jsoniter "github.com/json-iterator/go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
)

// UnaryServerLogInterceptor log 拦截
func UnaryServerLogInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	var addr string
	if peer, ok := peer.FromContext(ctx); ok {
		addr = peer.Addr.String()
	}

	var start = time.Now()
	defer func() {
		log.Printf("[I] addr: %v, cost: %v, service: %v, method: %v, req: %v, resp: %v, err: %v", addr, time.Since(start), path.Dir(info.FullMethod)[1:], path.Base(info.FullMethod), jsonFormat(req), jsonFormat(resp), err)
	}()

	resp, err = handler(ctx, req)
	return resp, err
}

// StreamServerRecoveryInterceptor recover
func StreamServerLogInterceptor(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) (err error) {
	var addr string
	if peer, ok := peer.FromContext(stream.Context()); ok {
		addr = peer.Addr.String()
	}
	var start = time.Now()
	defer func() {
		log.Printf("[I] addr: %v, cost: %v, service: %v, method: %v, srv: %v, err: %v", addr, time.Since(start), path.Dir(info.FullMethod)[1:], path.Base(info.FullMethod), jsonFormat(srv), err)
	}()

	return handler(srv, stream)
}

func jsonFormat(data interface{}) string {
	buf, err := jsoniter.ConfigCompatibleWithStandardLibrary.Marshal(data)
	if err == nil {
		return string(buf)
	}

	if a, ok := data.(StringAble); ok {
		return a.String()
	}

	return ""
}

// StringAble string
type StringAble interface {
	String() string
}
