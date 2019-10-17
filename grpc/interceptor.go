package gogrpc

import (
	"context"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"runtime/debug"
	"time"
)

// InterceptorUnary unary interceptor
var InterceptorUnary = func() grpc.ServerOption {
	interceptor := func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		// recover
		defer func() {
			if e := recover(); e != nil {
				err = errors.New(string(debug.Stack()))
				debug.PrintStack()
			}
		}()

		// auth
		ctx, err = Authorization(ctx, info.FullMethod)
		if err != nil {
			return nil, err
		}

		// log
		logReq := &RequestLogReq{
			Context: ctx,
			Method:  info.FullMethod,

			IsUnary:         true,    //unary
			ReqBody:         req,     //  body
			UnaryServerInfo: info,    // info
			UnaryHandler:    handler, // handler
		}

		// 现在时间
		nowTime := time.Now()

		// next
		resp, err = handler(ctx, req)
		if err != nil {
			// log error
			logReq.ExecError = err
			// wrap error
			err = errors.WithStack(err)
		}

		// log
		logReq.ExecDuration = time.Since(nowTime)
		go RequestLog(logReq)

		return resp, err
	}
	return grpc.UnaryInterceptor(interceptor)
}

// InterceptorStream stream interceptor
var InterceptorStream = func() grpc.ServerOption {
	var interceptor = func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) (err error) {
		// recover
		defer func() {
			if e := recover(); e != nil {
				err = errors.New(string(debug.Stack()))
				debug.PrintStack()
			}
		}()

		// auth
		ctx, err := Authorization(ss.Context(), info.FullMethod)
		if err != nil {
			return err
		}

		// log
		logReq := &RequestLogReq{
			Context: ctx,
			Method:  info.FullMethod,

			// stream server
			IsSteam:          true,    //stream
			Srv:              srv,     // srv
			ServerStream:     ss,      // stream
			StreamServerInfo: info,    // info
			StreamHandler:    handler, // handler
		}

		// 现在时间
		nowTime := time.Now()

		// next
		if err := handler(srv, ss); err != nil {
			// log error
			logReq.ExecError = err
			// wrap error
			err = errors.WithStack(err)
		}

		// log
		logReq.ExecDuration = time.Since(nowTime)
		// request log
		go RequestLog(logReq)

		return err
	}
	return grpc.StreamInterceptor(interceptor)
}
