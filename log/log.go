package golog

import (
	"github.com/pkg/errors"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

// default
const (
	defaultFilename       = "app.log"                // filename
	defaultRotationSuffix = ".%Y_%m_%d-%H_%M_%S.log" // rotation filename suffix
	defaultRotationMaxAge = -1                       // 100 years
)

// NewLog new log
func NewLog() (*logrus.Logger, error) {
	handler, err := newLog(InitConfig())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return handler, nil
}

// NewLogWithConfig new log
func NewLogWithConfig(cfg *Config) (*logrus.Logger, error) {
	handler, err := newLog(cfg)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return handler, nil
}

// newLog new log
func newLog(cfg *Config) (*logrus.Logger, error) {

	// handler
	handler := logrus.New()
	//handler.Formatter = new(logrus.JSONFormatter)
	handler.Level = logrus.DebugLevel

	// file system
	if err := RegisterFile(handler, cfg); err != nil {
		return nil, errors.WithStack(err)
	}

	// mysql
	if err := RegisterMysql(handler, cfg); err != nil {
		return nil, errors.WithStack(err)
	}
	return handler, nil
}

// RegisterMysql register mysql hook
func RegisterMysql(handler *logrus.Logger, cfg *Config) error {
	if !cfg.MysqlEnable {
		return nil
	}

	// mysql
	hook, err := MysqlHook(cfg)
	if err != nil {
		return errors.WithStack(err)
	}
	if hook != nil {
		AddHook(handler, hook)
	}
	return nil
}

// RegisterFile register file hook
func RegisterFile(handler *logrus.Logger, cfg *Config) error {
	if !cfg.FileEnable {
		return nil
	}

	// std || rotation
	switch {
	case cfg.FileRotation: // rotation
		hook, err := RotationHook(cfg)
		if err != nil {
			return errors.WithStack(err)
		}
		AddHook(handler, hook)
	default: // std
		hook, err := StdHook(cfg)
		if err != nil {
			return errors.WithStack(err)
		}
		AddHook(handler, hook)
	}
	return nil
}

// SetFormatter set formatter
func SetFormatter(logHandler *logrus.Logger, formatter logrus.Formatter) {
	logHandler.Formatter = formatter
}

// AddHook add hook
func AddHook(logHandler *logrus.Logger, hook logrus.Hook) {
	if hook == nil {
		return
	}
	logHandler.Hooks.Add(hook)
}

// WithField with field
func WithField(logHandler *logrus.Logger, key string, value interface{}) {
	logHandler.WithField(key, value)
}

// WithFields with fields
func WithFields(logHandler *logrus.Logger, fields map[string]interface{}) {
	if fields == nil {
		return
	}
	logHandler.WithFields(fields)
}

// StdHook std hook
var StdHook = func(cfg *Config) (*lfshook.LfsHook, error) {
	// filename
	if len(cfg.FileName) == 0 {
		cfg.FileName = defaultFilename
	}

	// abs path
	fPath := absPath(cfg.FileName)

	// path map
	var pathMap = lfshook.PathMap{
		logrus.PanicLevel: fPath,
		logrus.FatalLevel: fPath,
		logrus.ErrorLevel: fPath,
		logrus.WarnLevel:  fPath,
		logrus.InfoLevel:  fPath,
		logrus.DebugLevel: fPath,
	}
	return lfshook.NewHook(pathMap, new(logrus.JSONFormatter)), nil
}
