package godbconfigs

import (
	"github.com/pkg/errors"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

// database env
const (
	// connection
	EnvKeyDbDriver         = "AppDbDriver"         // driver
	EnvKeyDbEndpoints      = "AppDbEndpoints"      // host:post,host:post
	EnvKeyDbName           = "AppDbName"           // db name
	EnvKeyDbSetTablePrefix = "AppDbHasTablePrefix" // prefix
	EnvKeyDbTablePrefix    = "AppDbTablePrefix"    // prefix
	EnvKeyDbUser           = "AppDbUser"           // username
	EnvKeyDbPassword       = "AppDbPassword"       // password
	EnvKeyDbParameters     = "AppDbParameters"     // parameters

	// options
	EnvKeyDbDebug       = "AppDbDebug"       // debug
	EnvKeyDbMaxOpen     = "AppDbMaxOpen"     // max open
	EnvKeyDbMaxIdle     = "AppDbMaxIdle"     // idle
	EnvKeyDbMaxLifetime = "AppDbMaxLifetime" // lifetime
)

// Config config
type Config struct {
	Driver         string   `yaml:"driver"`           // os.Setenv(EnvKeyDbDriver, "mysql")
	Endpoints      []string `yaml:"endpoints"`        // os.Setenv(EnvKeyDbEndpoints, "127.0.0.1:3306,127.0.0.1:13306")
	DBName         string   `yaml:"db_name"`          // os.Setenv(EnvKeyDbName, "test")
	SetTablePrefix bool     `yaml:"set_table_prefix"` // os.Setenv(EnvKeyDbSetTablePrefix, "true")
	TablePrefix    string   `yaml:"table_prefix"`     // os.Setenv(EnvKeyDbTablePrefix, "ag_")
	Username       string   `yaml:"user"`             // os.Setenv(EnvKeyDbUser, "username")
	Password       string   `yaml:"password"`         // os.Setenv(EnvKeyDbPassword, "password")

	// parameters
	Parameters string `yaml:"parameters"` // os.Setenv(EnvKeyDbParameters, "charset=utf8&timeout=60s&loc=Local&autocommit=true")

	// sets orm LogMode
	Debug bool `yaml:"debug"` // os.Setenv(EnvKeyDbDebug, "true")

	// sets the maximum number of open connections to the database.
	MaxOpen int `yaml:"max_open"` // os.Setenv(EnvKeyDbMaxOpen, "10")

	// sets the maximum number of connections in the idle
	MaxIdle int `yaml:"max_idle"` // os.Setenv(EnvKeyDbMaxIdle, "10")

	// sets the maximum amount of time a connection may be reused.
	MaxLifetime time.Duration `yaml:"max_lifetime"` // os.Setenv(EnvKeyDbMaxLifetime, "30s")
}

// InitConfig init database config
var InitConfig = func() *Config {
	cfg := &Config{
		Driver:      strings.TrimSpace(os.Getenv(EnvKeyDbDriver)),
		Username:    strings.TrimSpace(os.Getenv(EnvKeyDbUser)),
		Password:    strings.TrimSpace(os.Getenv(EnvKeyDbPassword)),
		Endpoints:   strings.Split(strings.TrimSpace(os.Getenv(EnvKeyDbEndpoints)), ","),
		DBName:      strings.TrimSpace(os.Getenv(EnvKeyDbName)),
		Parameters:  strings.TrimSpace(os.Getenv(EnvKeyDbParameters)),
		TablePrefix: strings.TrimSpace(os.Getenv(EnvKeyDbTablePrefix)),
	}

	var err error

	// debug
	cfg.Debug, _ = strconv.ParseBool(strings.TrimSpace(os.Getenv(EnvKeyDbDebug)))

	// has table prefix
	if data := strings.TrimSpace(os.Getenv(EnvKeyDbSetTablePrefix)); len(data) > 0 {
		cfg.SetTablePrefix, err = strconv.ParseBool(data)
		if err != nil {
			log.Printf("%+v \n", errors.WithMessage(err, "strconv.Atoi(EnvKeyDbSetTablePrefix) fail"))
		}
	}

	// max open
	if data := strings.TrimSpace(os.Getenv(EnvKeyDbMaxOpen)); len(data) > 0 {
		cfg.MaxOpen, err = strconv.Atoi(data)
		if err != nil {
			log.Printf("%+v \n", errors.WithMessage(err, "strconv.Atoi(EnvKeyDbMaxOpen) fail"))
		}
	}

	// max idle
	if data := strings.TrimSpace(os.Getenv(EnvKeyDbMaxIdle)); len(data) > 0 {
		cfg.MaxIdle, err = strconv.Atoi(data)
		if err != nil {
			log.Printf("%+v \n", errors.WithMessage(err, "strconv.Atoi(EnvKeyDbMaxIdle) fail"))
		}
	}

	// max lifetime
	if data := strings.TrimSpace(os.Getenv(EnvKeyDbMaxLifetime)); len(data) > 0 {
		cfg.MaxLifetime, err = time.ParseDuration(data)
		if err != nil {
			log.Printf("%+v \n", errors.WithMessage(err, "time.ParseDuration(EnvKeyDbMaxLifetime) fail"))
		}
	}
	return cfg
}
