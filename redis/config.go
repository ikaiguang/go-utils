package goredis

import (
	"github.com/go-redis/redis"
	"github.com/pkg/errors"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

// env
const (
	EnvKeyRedisAddress            = "AppRedisAddress"            // host:port
	EnvKeyRedisPassword           = "AppRedisPassword"           // password
	EnvKeyRedisDB                 = "AppRedisDB"                 // DB
	EnvKeyRedisPoolSize           = "AppRedisPoolSize"           // pool
	EnvKeyRedisMinIdleConns       = "AppRedisMinIdleConns"       // min idle
	EnvKeyRedisMaxConnAge         = "AppRedisMaxConnAge"         // lifetime
	EnvKeyRedisDialTimeout        = "AppRedisDialTimeout"        // dial timeout
	EnvKeyRedisReadTimeout        = "AppRedisReadTimeout"        // read timeout
	EnvKeyRedisWriteTimeout       = "AppRedisWriteTimeout"       // write timeout
	EnvKeyRedisPoolTimeout        = "AppRedisPoolTimeout"        // pool timeout
	EnvKeyRedisIdleTimeout        = "AppRedisIdleTimeout"        // idle timeout
	EnvKeyRedisIdleCheckFrequency = "AppRedisIdleCheckFrequency" // idle check frequency
	EnvKeyRedisMaxRetries         = "AppRedisMaxRetries"         // retry
	EnvKeyRedisMinRetryBackoff    = "AppRedisMinRetryBackoff"    // min retry
	EnvKeyRedisMaxRetryBackoff    = "AppRedisMaxRetryBackoff"    // max retry
)

// InitConfig init config
var InitConfig = func() *redis.Options {
	var opt = &redis.Options{
		Addr:     strings.TrimSpace(os.Getenv(EnvKeyRedisAddress)),
		Password: strings.TrimSpace(os.Getenv(EnvKeyRedisPassword)),
	}

	var err error

	// Default is 10 connections per every CPU as reported by runtime.NumCPU.
	if data := strings.TrimSpace(os.Getenv(EnvKeyRedisPoolSize)); len(data) > 0 {
		opt.PoolSize, err = strconv.Atoi(data)
		if err != nil {
			log.Printf("%+v \n", errors.WithMessage(err, "strconv.Atoi(EnvKeyRedisPoolSize) fail"))
		}
	}

	// Minimum number of idle connections which is useful when establishing
	// new connection is slow.
	if data := strings.TrimSpace(os.Getenv(EnvKeyRedisMinIdleConns)); len(data) > 0 {
		opt.MinIdleConns, err = strconv.Atoi(data)
		if err != nil {
			log.Printf("%+v \n", errors.WithMessage(err, "strconv.Atoi(EnvKeyRedisMinIdleConns) fail"))
		}
	}

	// Default is to not close aged connections.
	if data := strings.TrimSpace(os.Getenv(EnvKeyRedisMaxConnAge)); len(data) > 0 {
		opt.MaxConnAge, err = time.ParseDuration(data)
		if err != nil {
			log.Printf("%+v \n", errors.WithMessage(err, "time.ParseDuration(EnvKeyRedisMaxConnAge) fail"))
		}
	}

	// Default is 5 seconds.
	if data := strings.TrimSpace(os.Getenv(EnvKeyRedisDialTimeout)); len(data) > 0 {
		opt.DialTimeout, err = time.ParseDuration(data)
		if err != nil {
			log.Printf("%+v \n", errors.WithMessage(err, "time.ParseDuration(EnvKeyRedisDialTimeout) fail"))
		}
	}

	// Default is 3 seconds.
	if data := strings.TrimSpace(os.Getenv(EnvKeyRedisReadTimeout)); len(data) > 0 {
		opt.ReadTimeout, err = time.ParseDuration(data)
		if err != nil {
			log.Printf("%+v \n", errors.WithMessage(err, "time.ParseDuration(EnvKeyRedisReadTimeout) fail"))
		}
	}

	// Default is ReadTimeout.
	if data := strings.TrimSpace(os.Getenv(EnvKeyRedisWriteTimeout)); len(data) > 0 {
		opt.WriteTimeout, err = time.ParseDuration(data)
		if err != nil {
			log.Printf("%+v \n", errors.WithMessage(err, "time.ParseDuration(EnvKeyRedisWriteTimeout) fail"))
		}
	}

	// Default is ReadTimeout + 1 second.
	if data := strings.TrimSpace(os.Getenv(EnvKeyRedisPoolTimeout)); len(data) > 0 {
		opt.PoolTimeout, err = time.ParseDuration(data)
		if err != nil {
			log.Printf("%+v \n", errors.WithMessage(err, "time.ParseDuration(EnvKeyRedisPoolTimeout) fail"))
		}
	}

	// Default is 5 minutes. -1 disables idle timeout check.
	if data := strings.TrimSpace(os.Getenv(EnvKeyRedisIdleTimeout)); len(data) > 0 {
		opt.IdleTimeout, err = time.ParseDuration(data)
		if err != nil {
			log.Printf("%+v \n", errors.WithMessage(err, "time.ParseDuration(EnvKeyRedisIdleTimeout) fail"))
		}
	}

	// Default is 1 minute. -1 disables idle connections reaper,
	if data := strings.TrimSpace(os.Getenv(EnvKeyRedisIdleCheckFrequency)); len(data) > 0 {
		opt.IdleCheckFrequency, err = time.ParseDuration(data)
		if err != nil {
			log.Printf("%+v \n", errors.WithMessage(err, "time.ParseDuration(EnvKeyRedisIdleCheckFrequency) fail"))
		}
	}

	// db
	if data := strings.TrimSpace(os.Getenv(EnvKeyRedisDB)); len(data) > 0 {
		opt.DB, err = strconv.Atoi(data)
		if err != nil {
			log.Printf("%+v \n", errors.WithMessage(err, "strconv.Atoi(EnvKeyRedisDB) fail"))
		}
	}

	// retry
	if data := strings.TrimSpace(os.Getenv(EnvKeyRedisMaxRetries)); len(data) > 0 {
		opt.MaxRetries, err = strconv.Atoi(data)
		if err != nil {
			log.Printf("%+v \n", errors.WithMessage(err, "strconv.Atoi(EnvKeyRedisMaxRetries) fail"))
		}
	}

	// Default is 8 milliseconds; -1 disables backoff.
	if data := strings.TrimSpace(os.Getenv(EnvKeyRedisMinRetryBackoff)); len(data) > 0 {
		opt.MinRetryBackoff, err = time.ParseDuration(data)
		if err != nil {
			log.Printf("%+v \n", errors.WithMessage(err, "time.ParseDuration(EnvKeyRedisMinRetryBackoff) fail"))
		}
	}

	//  Default is 512 milliseconds; -1 disables backoff.
	if data := strings.TrimSpace(os.Getenv(EnvKeyRedisMaxRetryBackoff)); len(data) > 0 {
		opt.MaxRetryBackoff, err = time.ParseDuration(data)
		if err != nil {
			log.Printf("%+v \n", errors.WithMessage(err, "time.ParseDuration(EnvKeyRedisMaxRetryBackoff) fail"))
		}
	}

	// SetDialer
	SetDialer(opt)

	// SetOnConnect
	SetOnConnect(opt)

	// SetTLS
	SetTLS(opt)

	return opt
}

// SetDialer creates new network connection and has priority over Network and Addr options.
func SetDialer(opt *redis.Options) {
	//opt.Dialer = func() (net.Conn, error) { return nil, nil }
}

// SetOnConnect Hook that is called when new connection is established.
func SetOnConnect(opt *redis.Options) {
	//opt.OnConnect = func(*redis.Conn) error { return nil }
}

// SetTLS tls.Config
func SetTLS(opt *redis.Options) {
	//opt.TLSConfig = new(tls.Config)
}
