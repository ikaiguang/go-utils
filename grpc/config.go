package gogrpc

import (
	"github.com/pkg/errors"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// server config
const (
	defaultAddress = ":50051" // address
)

// server mode
const (
	DebugMode   = "debug"   // debug.
	ReleaseMode = "release" // release.
	TestMode    = "test"    // test.
)

// server env
const (
	EnvKeyRpcServerAddr    = "AppRpcServerAddress" // server address
	EnvKeyRpcServerMode    = "AppRpcServerMode"    // server model
	EnvKeyRpcSSLEnable     = "AppRpcSSLEnable"     // ssl enable
	EnvKeyRpcSSLServerName = "AppRpcSSLServerName" // ssl server name
	EnvKeyRpcSSLServerCrt  = "AppRpcSSLServerCrt"  // ssl server crt
	EnvKeyRpcSSLServerKey  = "AppRpcSSLServerKey"  // ssl server key
	EnvKeyRpcSSLClientCrt  = "AppRpcSSLClientCrt"  // ssl client crt
	EnvKeyRpcRegAddress    = "AppRpcRegAddress"    // reg address
	EnvKeyRpcRegDomain     = "AppRpcRegDomain"     // reg domain
	EnvKeyRpcOssDomain     = "AppRpcOssDomain"     // oss domain
)

// config
var (
	cfg *Config // config
)

// getConfig new config
func getConfig() *Config {
	if cfg != nil {
		return cfg
	}
	return newConfig()
}

// newConfig new config
func newConfig() *Config {
	cfg = InitConfig()

	// address
	if len(cfg.Address) == 0 {
		cfg.Address = defaultAddress
	}

	return cfg
}

// Config server config
type Config struct {
	Address       string `yaml:"server_address"`  // server address
	Mode          string `yaml:"server_mode"`     // server mode
	SSLEnable     bool   `yaml:"ssl_enable"`      // ssl enable
	SSLServerName string `yaml:"ssl_server_name"` // ssl server name
	SSLServerKey  string `yaml:"ssl_server_key"`  // ssl server key
	SSLServerCrt  string `yaml:"ssl_server_crt"`  // ssl server crt
	SSLClientCrt  string `yaml:"ssl_client_crt"`  // ssl client crt
	RegAddress    string `yaml:"reg_address"`     // reg address
	RegDomain     string `yaml:"reg_domain"`      // reg domain
	OssDomain     string `yaml:"oss_domain"`      // oss domain
}

// InitConfig init config
var InitConfig = func() *Config {
	var cfg = &Config{
		Address:       strings.TrimSpace(os.Getenv(EnvKeyRpcServerAddr)),
		Mode:          strings.TrimSpace(os.Getenv(EnvKeyRpcServerMode)),
		SSLEnable:     false,
		SSLServerName: strings.TrimSpace(os.Getenv(EnvKeyRpcSSLServerName)),
		SSLServerCrt:  strings.TrimSpace(os.Getenv(EnvKeyRpcSSLServerCrt)),
		SSLServerKey:  strings.TrimSpace(os.Getenv(EnvKeyRpcSSLServerKey)),
		SSLClientCrt:  strings.TrimSpace(os.Getenv(EnvKeyRpcSSLClientCrt)),
		RegAddress:    strings.TrimSpace(os.Getenv(EnvKeyRpcRegAddress)),
		RegDomain:     strings.TrimSpace(os.Getenv(EnvKeyRpcRegDomain)),
		OssDomain:     strings.TrimSpace(os.Getenv(EnvKeyRpcOssDomain)),
	}

	// mode
	switch cfg.Mode {
	case ReleaseMode, DebugMode, TestMode:
	default:
		cfg.Mode = ReleaseMode
	}

	// ssl enable
	cfg.SSLEnable, _ = strconv.ParseBool(strings.TrimSpace(os.Getenv(EnvKeyRpcSSLEnable)))

	// file path
	pwdPath, err := os.Getwd()
	if err != nil {
		log.Printf("%+v \n", errors.WithStack(err))
		return cfg
	}

	// server crt
	if len(cfg.SSLServerCrt) > 0 && !filepath.IsAbs(cfg.SSLServerCrt) {
		cfg.SSLServerCrt = filepath.Join(pwdPath, cfg.SSLServerCrt)
	}

	// server key
	if len(cfg.SSLServerKey) > 0 && !filepath.IsAbs(cfg.SSLServerKey) {
		cfg.SSLServerKey = filepath.Join(pwdPath, cfg.SSLServerKey)
	}

	// client crt
	if len(cfg.SSLClientCrt) > 0 && !filepath.IsAbs(cfg.SSLClientCrt) {
		cfg.SSLClientCrt = filepath.Join(pwdPath, cfg.SSLClientCrt)
	}
	return cfg
}
