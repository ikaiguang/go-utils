package goetcd

import (
	"github.com/coreos/etcd/clientv3"
	"github.com/pkg/errors"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

// etcd config
const (
	defaultETCDDialTimeout = time.Second * 5 // dial timeout
)

// etcd env
const (
	EnvKeyETCDEndPoints            = "AppETCDEndpoints"                // endpoints
	EnvKeyETCDUsername             = "AppETCDUsername"                 // username
	EnvKeyETCDPassword             = "AppETCDPassword"                 // password
	EnvKeyETCDDialTimeout          = "AppETCDDialTimeout"              // timeout
	EnvKeyETCDDialKeepAliveTime    = "AppETCDDialKeepAliveTime"        // keepalive
	EnvKeyETCDDialKeepAliveTimeout = "AppETCDDialKeepAliveTimeTimeout" // keepalive timeout
	EnvKeyETCDAutoSyncInterval     = "AppETCDAutoSyncInterval"         // default auto-sync is disabled.
	EnvKeyETCDMaxCallSendMsgSize   = "AppETCDMaxCallSendMsgSize"       // If 0, it defaults to 2.0 MiB (2 * 1024 * 1024)
	EnvKeyETCDMaxCallRecvMsgSize   = "AppETCDMaxCallRecvMsgSize"       // If 0, it defaults to "math.MaxInt32"
	EnvKeyETCDRejectOldCluster     = "AppETCDRejectOldCluster"         // when set will refuse to create a client against an outdated cluster.
)

// InitConfig init config
var InitConfig = func() *clientv3.Config {
	// config
	var cfg = &clientv3.Config{
		Endpoints: strings.Split(strings.TrimSpace(os.Getenv(EnvKeyETCDEndPoints)), ","),
		Username:  strings.TrimSpace(os.Getenv(EnvKeyETCDUsername)),
		Password:  strings.TrimSpace(os.Getenv(EnvKeyETCDPassword)),
	}

	var err error

	// timeout
	if data := strings.TrimSpace(os.Getenv(EnvKeyETCDDialTimeout)); len(data) > 0 {
		cfg.DialTimeout, err = time.ParseDuration(data)
		if err != nil {
			log.Printf("%+v \n", errors.WithMessage(err, "time.ParseDuration(EnvKeyETCDDialTimeout) fail"))
		}
	}

	// keepalive
	if data := strings.TrimSpace(os.Getenv(EnvKeyETCDDialKeepAliveTime)); len(data) > 0 {
		cfg.DialKeepAliveTime, err = time.ParseDuration(data)
		if err != nil {
			log.Printf("%+v \n", errors.WithMessage(err, "time.ParseDuration(EnvKeyETCDDialKeepAliveTime) fail"))
		}
	}

	// keepalive timeout
	if data := strings.TrimSpace(os.Getenv(EnvKeyETCDDialKeepAliveTimeout)); len(data) > 0 {
		cfg.DialKeepAliveTimeout, err = time.ParseDuration(data)
		if err != nil {
			log.Printf("%+v \n", errors.WithMessage(err, "time.ParseDuration(EnvKeyETCDDialKeepAliveTimeout) fail"))
		}
	}

	// auto-sync
	if data := strings.TrimSpace(os.Getenv(EnvKeyETCDAutoSyncInterval)); len(data) > 0 {
		cfg.AutoSyncInterval, err = time.ParseDuration(data)
		if err != nil {
			log.Printf("%+v \n", errors.WithMessage(err, "time.ParseDuration(EnvKeyETCDAutoSyncInterval) fail"))
		}
	}

	// MaxCallSendMsgSize is the client-side request send limit in bytes.
	// If 0, it defaults to 2.0 MiB (2 * 1024 * 1024).
	if data := strings.TrimSpace(os.Getenv(EnvKeyETCDMaxCallSendMsgSize)); len(data) > 0 {
		cfg.MaxCallSendMsgSize, err = strconv.Atoi(data)
		if err != nil {
			log.Printf("%+v \n", errors.WithMessage(err, "strconv.Atoi(EnvKeyETCDMaxCallSendMsgSize) fail"))
		}
	}

	// MaxCallRecvMsgSize is the client-side response receive limit.
	// If 0, it defaults to "math.MaxInt32", because range response can easily exceed request send limits.
	if data := strings.TrimSpace(os.Getenv(EnvKeyETCDMaxCallRecvMsgSize)); len(data) > 0 {
		cfg.MaxCallRecvMsgSize, err = strconv.Atoi(data)
		if err != nil {
			log.Printf("%+v \n", errors.WithMessage(err, "strconv.Atoi(EnvKeyETCDMaxCallRecvMsgSize) fail"))
		}
	}

	// RejectOldCluster when set will refuse to create a client against an outdated cluster.
	if data := strings.TrimSpace(os.Getenv(EnvKeyETCDRejectOldCluster)); len(data) > 0 {
		cfg.RejectOldCluster, err = strconv.ParseBool(data)
		if err != nil {
			log.Printf("%+v \n", errors.WithMessage(err, "strconv.ParseBool(EnvKeyETCDRejectOldCluster) fail"))
		}
	}

	// context
	SetContext(cfg)

	// tls
	SetTLS(cfg)

	// dial options
	AppendDialOptions(cfg)

	return cfg
}

// SetContext config context
var SetContext = func(cfg *clientv3.Config) {
	//cfg.Context = context.Background()
}

// SetTLS config tls
var SetTLS = func(cfg *clientv3.Config) {
	//cfg.TLS = new(tls.Config)
}

// AppendDialOptions config tls
var AppendDialOptions = func(cfg *clientv3.Config) {
	//cfg.DialOptions = append(cfg.DialOptions, []grpc.DialOption{}...)
}
