package gogrpc

import (
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"net"
)

// NewServer grpc server
func NewServer() (*grpc.Server, error) {
	return newServer(getConfig())
}

// NewServerWithConfig grpc server
func NewServerWithConfig(cfg *Config) (*grpc.Server, error) {
	return newServer(cfg)
}

// newServer grpc server
func newServer(cfg *Config) (*grpc.Server, error) {
	// options
	var opts []grpc.ServerOption

	// ssl
	if cfg.SSLEnable {
		cred, err := credentials.NewServerTLSFromFile(cfg.SSLServerCrt, cfg.SSLServerKey)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		opts = append(opts, grpc.Creds(cred))
	}

	// unary interceptor
	opts = append(opts, InterceptorUnary())

	// stream interceptor
	opts = append(opts, InterceptorStream())

	// other options
	options, err := SetServerOptions()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	opts = append(opts, options...)
	return grpc.NewServer(opts...), nil
}

// SetServerOptions server options
var SetServerOptions = func() ([]grpc.ServerOption, error) {
	var opts []grpc.ServerOption

	// max message size default 4M
	//opts = append(opts, grpc.MaxRecvMsgSize(15*1024*1024))

	return opts, nil
}

// RunServer start
func RunServer(server *grpc.Server) error {
	return runServer(server, getConfig())
}

// RunServerWithConfig start
func RunServerWithConfig(server *grpc.Server, cfg *Config) error {
	return runServer(server, cfg)
}

// runServer start
func runServer(server *grpc.Server, cfg *Config) error {
	// tcp
	lis, err := net.Listen("tcp", cfg.Address)
	if err != nil {
		return errors.WithStack(err)
	}

	// info
	log.Printf("server addr(%s), isSSL(%v) \n", cfg.Address, cfg.SSLEnable)

	// register
	if err := RegisterServer(server); err != nil {
		return errors.WithStack(err)
	}

	// start
	if err := server.Serve(lis); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// RunServer start
var RegisterServer = func(server *grpc.Server) error {
	return nil
}
