package golog

import (
	"github.com/lestrrat-go/file-rotatelogs"
	"github.com/pkg/errors"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

// RotationHook rotation hook
var RotationHook = func(cfg *Config) (*lfshook.LfsHook, error) {
	// filename
	if len(cfg.FileName) == 0 {
		cfg.FileName = defaultFilename
	}

	// suffix
	if len(cfg.FileOptRotationSuffix) == 0 {
		cfg.FileOptRotationSuffix = defaultRotationSuffix
	}

	// abs path
	fPath := absPath(cfg.FileName) + cfg.FileOptRotationSuffix

	// option
	var opts []rotatelogs.Option

	// time location
	if cfg.FileOptTimeLocation != nil {
		opts = append(opts, rotatelogs.WithLocation(cfg.FileOptTimeLocation))
	}

	// link name
	if len(cfg.FileOptLinkName) > 0 {
		opts = append(opts, rotatelogs.WithLinkName(cfg.FileOptLinkName))
	}

	// force new file
	if cfg.FileOptForceNewFile {
		opts = append(opts, rotatelogs.ForceNewFile())
	}

	// lifetime
	if cfg.FileOptMaxAge > 0 {
		opts = append(opts, rotatelogs.WithMaxAge(cfg.FileOptMaxAge))
	} else {
		opts = append(opts, rotatelogs.WithMaxAge(defaultRotationMaxAge))
	}

	// rotation time
	if cfg.FileOptRotationTime > 0 {
		opts = append(opts, rotatelogs.WithRotationTime(cfg.FileOptRotationTime))
	}

	// rotation count
	if cfg.FileOptRotationCount > 0 {
		opts = append(opts, rotatelogs.WithRotationCount(cfg.FileOptRotationCount))
	}

	// handler
	opts = append(opts, WithHandler(cfg)...)

	// writer
	writer, err := rotatelogs.New(fPath, opts...)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	// writer map
	var writerMap = lfshook.WriterMap{
		logrus.PanicLevel: writer,
		logrus.FatalLevel: writer,
		logrus.ErrorLevel: writer,
		logrus.WarnLevel:  writer,
		logrus.InfoLevel:  writer,
		logrus.DebugLevel: writer,
	}
	return lfshook.NewHook(writerMap, new(logrus.JSONFormatter)), nil
}

// WithHandler with handler
var WithHandler = func(cfg *Config) []rotatelogs.Option {
	//cfg.FileOpts = append(cfg.FileOpts, rotatelogs.WithHandler(h))
	return nil
}
