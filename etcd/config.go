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
	EnvKeyETCDEndPoints            = "AppETCDEndpoints"                // endpoints:host:post,host:post
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

// Config config
type Config struct {
	Endpoints            []string      `yaml:"endpoints"`               // os.Setenv(EnvKeyETCDEndPoints, "127.0.0.1:2379,localhost:2379")
	Username             string        `yaml:"username"`                // os.Setenv(AppETCDUsername, "")
	Password             string        `yaml:"password"`                // os.Setenv(AppETCDPassword, "")
	DialTimeout          time.Duration `yaml:"dial_timeout"`            // os.Setenv(AppETCDDialTimeout, "3s")
	DialKeepAliveTime    time.Duration `yaml:"dial_keep_alive_time"`    // os.Setenv(AppETCDDialKeepAliveTime, "0s")
	DialKeepAliveTimeout time.Duration `yaml:"dial_keep_alive_timeout"` // os.Setenv(AppETCDDialKeepAliveTimeTimeout, "0s")
	AutoSyncInterval     time.Duration `yaml:"auto_sync_interval"`      // os.Setenv(AppETCDAutoSyncInterval, "0s")
	MaxCallSendMsgSize   int           `yaml:"max_call_send_msg_size"`  // os.Setenv(AppETCDMaxCallSendMsgSize, "0")
	MaxCallRecvMsgSize   int           `yaml:"max_call_recv_msg_size"`  // os.Setenv(AppETCDMaxCallRecvMsgSize, "0")
	RejectOldCluster     bool          `yaml:"reject_old_cluster"`      // os.Setenv(AppETCDRejectOldCluster, "0")
}

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

// InitConfigFromCustom init config
var InitConfigFromCustom = func(custom *Config) *clientv3.Config {
	// config
	var cfg = &clientv3.Config{
		Endpoints:            custom.Endpoints,
		Username:             custom.Username,
		Password:             custom.Password,
		DialTimeout:          custom.DialTimeout,
		DialKeepAliveTime:    custom.DialKeepAliveTime,
		DialKeepAliveTimeout: custom.DialKeepAliveTimeout,
		AutoSyncInterval:     custom.AutoSyncInterval,
		MaxCallSendMsgSize:   custom.MaxCallSendMsgSize,
		MaxCallRecvMsgSize:   custom.MaxCallRecvMsgSize,
		RejectOldCluster:     custom.RejectOldCluster,
	}

	// context
	SetContext(cfg)

	// tls
	SetTLS(cfg)

	// dial options
	AppendDialOptions(cfg)

	return cfg
}
