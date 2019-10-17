package gogrpc

import (
	"context"
)

// Authorization grpc auth
// @Param fullMethod grpc.serverInfo.FullMethod(grpc.UnaryServerInfo || info *grpc.StreamServerInfo)
var Authorization = func(ctx context.Context, fullMethod string) (context.Context, error) {
	// metadata
	//if md, ok := metadata.FromIncomingContext(ctx); ok {
	//	// oauth2 authorization
	//	if data, ok := md["authorization"]; ok {
	//		_ = data
	//	}
	//}
	return ctx, nil
}
