package gohttp

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// server config
const (
	defaultAddress = ":8999" // address
)

// server env
const (
	EnvKeyHttpServerAddr    = "AppHttpServerAddress" // server address
	EnvKeyHttpServerMode    = "AppHttpServerMode"    // server model
	EnvKeyHttpSSLEnable     = "AppHttpSSLEnable"     // ssl enable
	EnvKeyHttpSSLServerName = "AppHttpSSLServerName" // ssl server name
	EnvKeyHttpSSLServerCrt  = "AppHttpSSLServerCrt"  // ssl server crt
	EnvKeyHttpSSLServerKey  = "AppHttpSSLServerKey"  // ssl server key
	EnvKeyHttpSSLClientCrt  = "AppHttpSSLClientCrt"  // ssl client crt
	EnvKeyHttpRegAddress    = "AppHttpRegAddress"    // reg address
	EnvKeyHttpRegDomain     = "AppHttpRegDomain"     // reg domain
	EnvKeyHttpOssDomain     = "AppHttpOssDomain"     // oss domain
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
		Address:       strings.TrimSpace(os.Getenv(EnvKeyHttpServerAddr)),
		Mode:          strings.TrimSpace(os.Getenv(EnvKeyHttpServerMode)),
		SSLEnable:     false,
		SSLServerName: strings.TrimSpace(os.Getenv(EnvKeyHttpSSLServerName)),
		SSLServerCrt:  strings.TrimSpace(os.Getenv(EnvKeyHttpSSLServerCrt)),
		SSLServerKey:  strings.TrimSpace(os.Getenv(EnvKeyHttpSSLServerKey)),
		SSLClientCrt:  strings.TrimSpace(os.Getenv(EnvKeyHttpSSLClientCrt)),
		RegAddress:    strings.TrimSpace(os.Getenv(EnvKeyHttpRegAddress)),
		RegDomain:     strings.TrimSpace(os.Getenv(EnvKeyHttpRegDomain)),
		OssDomain:     strings.TrimSpace(os.Getenv(EnvKeyHttpOssDomain)),
	}

	// mode
	switch cfg.Mode {
	case gin.ReleaseMode, gin.DebugMode, gin.TestMode:
	default:
		cfg.Mode = gin.ReleaseMode
	}

	// ssl enable
	cfg.SSLEnable, _ = strconv.ParseBool(strings.TrimSpace(os.Getenv(EnvKeyHttpSSLEnable)))

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
