package goetcd

import (
	"context"
	"github.com/coreos/etcd/clientv3"
	"github.com/pkg/errors"
)

// NewClient new client
func NewClient() (*clientv3.Client, error) {
	return newClient(InitConfig())
}

// NewClientWithConfig new client
func NewClientWithConfig(cfg *clientv3.Config) (*clientv3.Client, error) {
	return newClient(cfg)
}

// newClient new client
func newClient(cfg *clientv3.Config) (*clientv3.Client, error) {
	// timeout
	if cfg.DialTimeout == 0 {
		cfg.DialTimeout = defaultETCDDialTimeout
	}

	// client
	eClient, err := clientv3.New(*cfg)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	// ping
	ctx, cancel := context.WithTimeout(context.Background(), defaultETCDDialTimeout)
	defer cancel()
	if _, err := eClient.Grant(ctx, 0); err != nil {
		return nil, errors.WithStack(err)
	}
	return eClient, nil
}
