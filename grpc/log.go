package gogrpc

import (
	"context"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"runtime/debug"
	"time"
)

// RequestLogReq log param
type RequestLogReq struct {
	Context context.Context // context
	Method  string          //  method

	// unary server
	IsUnary         bool                  // unary
	ReqBody         interface{}           //  body
	UnaryServerInfo *grpc.UnaryServerInfo // info
	UnaryHandler    grpc.UnaryHandler     // handler

	// stream server
	IsSteam          bool                   // stream
	Srv              interface{}            // srv
	ServerStream     grpc.ServerStream      // stream
	StreamServerInfo *grpc.StreamServerInfo // info
	StreamHandler    grpc.StreamHandler     // handler

	// response
	ExecDuration time.Duration // latency
	ExecError    error         // error
}

// RequestLog 请求日志
var RequestLog = func(req *RequestLogReq) (err error) {
	// recover
	defer func() {
		if e := recover(); e != nil {
			err = errors.New(string(debug.Stack()))
			debug.PrintStack()
		}
	}()
	// metadata
	//if md, ok := metadata.FromIncomingContext(req.Context); ok {
	//	// authorization
	//	if data, ok := md["authorization"]; ok {
	//		_ = data
	//	}
	//}
	switch {
	case req.IsUnary: // unary
	case req.IsSteam: // steam

	}
	return err
}
