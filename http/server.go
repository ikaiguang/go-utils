package gohttp

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"log"
)

// NewServer new server
func NewServer() *gin.Engine {
	return newServer(getConfig())
}

// NewServerWithConfig new server
func NewServerWithConfig(cfg *Config) *gin.Engine {
	return newServer(cfg)
}

// newServer new server
func newServer(cfg *Config) *gin.Engine {
	// mode
	gin.SetMode(cfg.Mode)

	// server
	if cfg.Mode == gin.ReleaseMode {
		return gin.New()
	}
	return gin.Default()
}

// RunServer run server
func RunServer(engine *gin.Engine) error {
	return runServer(engine, getConfig())
}

// RunServerWithConfig run server
func RunServerWithConfig(engine *gin.Engine, cfg *Config) error {
	return runServer(engine, cfg)
}

// runServer run server
func runServer(engine *gin.Engine, cfg *Config) error {
	// info
	log.Printf("server addr(%s), isSSL(%v) \n", cfg.Address, cfg.SSLEnable)

	// ssl
	if cfg.SSLEnable {
		if err := engine.RunTLS(cfg.Address, cfg.SSLServerCrt, cfg.SSLServerKey); err != nil {
			return errors.WithStack(err)
		}
	}

	// start
	if err := engine.Run(cfg.Address); err != nil {
		return errors.WithStack(err)
	}
	return nil
}
