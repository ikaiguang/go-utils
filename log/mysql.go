package golog

import (
	"fmt"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

// mysqlWriter mysql writer
type mysqlWriter int

func (w *mysqlWriter) Write(p []byte) (n int, err error) {
	fmt.Printf("mysql writer : message len(%d) \n", len(p))
	fmt.Println("catch message : ", string(p))
	return len(p), nil
}

// MysqlHook mysql hook
var MysqlHook = func(cfg *Config) (*lfshook.LfsHook, error) {
	writer := new(mysqlWriter)

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
