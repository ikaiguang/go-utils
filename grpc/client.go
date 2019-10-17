package gogrpc

import (
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
)

// NewClient new client
func NewClient() (*grpc.ClientConn, error) {
	return newClient(getConfig())
}

// NewClientWithConfig new client
func NewClientWithConfig(cfg *Config) (*grpc.ClientConn, error) {
	return newClient(cfg)
}

// newClient new client
func newClient(cfg *Config) (*grpc.ClientConn, error) {
	// options
	opts, err := GetClientOptions(cfg)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	// client
	conn, err := grpc.Dial(cfg.Address, opts...)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	// info
	log.Printf("dial addr : %s \n", cfg.Address)

	return conn, nil
}

// NewClientWithAddr new client
func NewClientWithAddr(address string) (*grpc.ClientConn, error) {
	return newClientWithAddr(getConfig(), address)
}

// NewClientWithConfigAndAddr new client
func NewClientWithConfigAndAddr(cfg *Config, address string) (*grpc.ClientConn, error) {
	return newClientWithAddr(cfg, address)
}

// newClientWithAddr new client
func newClientWithAddr(cfg *Config, address string) (*grpc.ClientConn, error) {
	// options
	opts, err := GetClientOptions(cfg)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	// address
	if len(address) == 0 {
		address = cfg.Address
	}

	// client
	conn, err := grpc.Dial(address, opts...)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	// info
	log.Printf("dial addr : %s \n", cfg.Address)

	return conn, nil
}

// NewClientWithResolver balancer resolver
var NewClientWithResolver = func() (*grpc.ClientConn, error) {
	return nil, nil
}

// GetClientOptions client option
func GetClientOptions(cfg *Config) ([]grpc.DialOption, error) {
	// options
	var opts []grpc.DialOption
	// ssl
	if cfg.SSLEnable {
		cred, err := credentials.NewClientTLSFromFile(cfg.SSLClientCrt, cfg.SSLServerName)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		opts = append(opts, grpc.WithTransportCredentials(cred))
	} else {
		opts = append(opts, grpc.WithInsecure())
	}

	// other options
	options, err := SetClientOptions()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	opts = append(opts, options...)
	return opts, nil
}

// SetClientOptions client options
var SetClientOptions = func() ([]grpc.DialOption, error) {
	var opts []grpc.DialOption

	// option
	//var optsCallOption = []grpc.CallOption{
	//	grpc.MaxCallRecvMsgSize(15 * 1024 * 1024),
	//	grpc.MaxCallSendMsgSize(15 * 1024 * 1024),
	//}
	//opts = append(opts, grpc.WithDefaultCallOptions(optsCallOption...))

	return opts, nil
}
