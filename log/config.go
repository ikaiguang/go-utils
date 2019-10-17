package golog

import (
	"github.com/pkg/errors"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

// server env
const (
	EnvKeyLogMysqlEnable           = "AppLogMysqlEnable"           // mysql enable
	EnvKeyLogFileEnable            = "AppLogFileEnable"            // file system enable
	EnvKeyLogFileRotation          = "AppLogFileRotation"          // file rotation
	EnvKeyLogFilename              = "AppLogFilename"              // filename
	EnvKeyLogFileOptTimeLocation   = "AppLogFileOptTimeLocal"      // time location
	EnvKeyLogFileOptLinkName       = "AppLogFileOptLinkName"       // link name
	EnvKeyLogFileOptForceNewFile   = "AppLogFileOptForceNewFile"   // force new file
	EnvKeyLogFileOptMaxAge         = "AppLogFileOptMaxAge"         // lifetime
	EnvKeyLogFileOptRotationTime   = "AppLogFileOptRotationTime"   // rotation time
	EnvKeyLogFileOptRotationCount  = "AppLogFileOptRotationCount"  // rotation count
	EnvKeyLogFileOptRotationSuffix = "AppLogFileOptRotationSuffix" // rotation suffix(example: ".%Y_%m_%d-%H_%M_%S.log")
)

// Config config
type Config struct {
	MysqlEnable           bool           // mysql
	FileEnable            bool           // file system
	FileRotation          bool           // file rotation
	FileName              string         // filename
	FileOptTimeLocation   *time.Location // default local (default: rotatelogs.Local)
	FileOptLinkName       string         // link name (default: "")
	FileOptForceNewFile   bool           // force new file (default: false)
	FileOptMaxAge         time.Duration  // lifetime (default: 7 days)
	FileOptRotationTime   time.Duration  // rotation time(default: 86400 sec)
	FileOptRotationCount  uint           // rotation count (default: -1)
	FileOptRotationSuffix string         // rotation suffix(example:path+".%Y_%m_%d-%H_%M_%S.log")
}

// InitConfig init config
var InitConfig = func() *Config {
	var cfg = &Config{
		FileName:              strings.TrimSpace(os.Getenv(EnvKeyLogFilename)),
		FileOptLinkName:       strings.TrimSpace(os.Getenv(EnvKeyLogFileOptLinkName)),
		FileOptRotationSuffix: strings.TrimSpace(os.Getenv(EnvKeyLogFileOptRotationSuffix)),
	}

	var err error

	// mysql enable
	if data := strings.TrimSpace(os.Getenv(EnvKeyLogMysqlEnable)); len(data) > 0 {
		cfg.MysqlEnable, err = strconv.ParseBool(data)
		if err != nil {
			log.Printf("%+v \n", errors.WithMessage(err, "strconv.ParseBool(EnvKeyLogMysqlEnable) fail"))
		}
	}

	// file enable
	if data := strings.TrimSpace(os.Getenv(EnvKeyLogFileEnable)); len(data) > 0 {
		cfg.FileEnable, err = strconv.ParseBool(data)
		if err != nil {
			log.Printf("%+v \n", errors.WithMessage(err, "strconv.ParseBool(EnvKeyLogFileEnable) fail"))
		}
	}

	// file rotation
	if data := strings.TrimSpace(os.Getenv(EnvKeyLogFileRotation)); len(data) > 0 {
		cfg.FileRotation, err = strconv.ParseBool(data)
		if err != nil {
			log.Printf("%+v \n", errors.WithMessage(err, "strconv.ParseBool(EnvKeyLogFileRotation) fail"))
		}
	}

	// time location
	if data := strings.TrimSpace(os.Getenv(EnvKeyLogFileOptTimeLocation)); len(data) > 0 {
		cfg.FileOptTimeLocation, err = time.LoadLocation(data)
		if err != nil {
			log.Printf("%+v \n", errors.WithMessage(err, "time.LoadLocation(EnvKeyLogFileOptTimeLocation) fail"))
		}
	}

	// force new file
	if data := strings.TrimSpace(os.Getenv(EnvKeyLogFileOptForceNewFile)); len(data) > 0 {
		cfg.FileOptForceNewFile, err = strconv.ParseBool(data)
		if err != nil {
			log.Printf("%+v \n", errors.WithMessage(err, "strconv.ParseBool(EnvKeyLogFileOptForceNewFile) fail"))
		}
	}

	// lifetime
	if data := strings.TrimSpace(os.Getenv(EnvKeyLogFileOptMaxAge)); len(data) > 0 {
		cfg.FileOptMaxAge, err = time.ParseDuration(data)
		if err != nil {
			log.Printf("%+v \n", errors.WithMessage(err, "time.ParseDuration(EnvKeyLogFileOptMaxAge) fail"))
		}
	}

	// rotation time
	if data := strings.TrimSpace(os.Getenv(EnvKeyLogFileOptRotationTime)); len(data) > 0 {
		cfg.FileOptRotationTime, err = time.ParseDuration(data)
		if err != nil {
			log.Printf("%+v \n", errors.WithMessage(err, "time.ParseDuration(EnvKeyLogFileOptRotationTime) fail"))
		}
	}

	// rotation count
	if data := strings.TrimSpace(os.Getenv(EnvKeyLogFileOptRotationCount)); len(data) > 0 {
		if rotationCount, err := strconv.Atoi(data); err != nil {
			log.Printf("%+v \n", errors.WithMessage(err, "strconv.Atoi(EnvKeyLogFileOptRotationCount) fail"))
		} else if rotationCount > 0 {
			cfg.FileOptRotationCount = uint(rotationCount)
		}
	}
	return cfg
}

// pwd pwd
var pwd, _ = os.Getwd()

// absPath abs path
func absPath(path string) string {
	if !filepath.IsAbs(path) {
		return filepath.Join(pwd, path)
	}
	return path
}
