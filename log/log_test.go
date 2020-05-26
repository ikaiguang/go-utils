package golog

import (
	"os"
	"testing"
)

func TestNewLog(t *testing.T) {
	os.Setenv(EnvKeyLogMysqlEnable, "false")
	os.Setenv(EnvKeyLogFileEnable, "true")
	os.Setenv(EnvKeyLogFileRotation, "true")
	os.Setenv(EnvKeyLogFilename, "_output/test.log")
	os.Setenv(EnvKeyLogFileOptTimeLocation, "Local")
	os.Setenv(EnvKeyLogFileOptLinkName, "")
	os.Setenv(EnvKeyLogFileOptForceNewFile, "true")
	os.Setenv(EnvKeyLogFileOptMaxAge, "0")
	os.Setenv(EnvKeyLogFileOptRotationTime, "0s")
	os.Setenv(EnvKeyLogFileOptRotationCount, "0")
	os.Setenv(EnvKeyLogFileOptRotationSuffix, ".%Y_%m_%d.log")

	handler, err := NewLog()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	handler.Info("info message")

	//if _, err := os.Open("./test.log"); err != nil {
	//	t.Error(err)
	//	t.FailNow()
	//}
}

func TestNewLog_Rotation(t *testing.T) {
	os.Setenv(EnvKeyLogMysqlEnable, "true")
	os.Setenv(EnvKeyLogFileEnable, "true")
	os.Setenv(EnvKeyLogFileRotation, "true")
	os.Setenv(EnvKeyLogFilename, "_output/test.log")
	os.Setenv(EnvKeyLogFileOptTimeLocation, "Local")
	os.Setenv(EnvKeyLogFileOptLinkName, "")
	os.Setenv(EnvKeyLogFileOptForceNewFile, "true")
	os.Setenv(EnvKeyLogFileOptMaxAge, "0")
	os.Setenv(EnvKeyLogFileOptRotationTime, "0s")
	os.Setenv(EnvKeyLogFileOptRotationCount, "0")
	os.Setenv(EnvKeyLogFileOptRotationSuffix, ".%Y_%m_%d.log")

	handler, err := NewLog()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	handler.Info("info message")

	//logFile := "./test.log." + time.Now().Format("2006_01_02")
	//if _, err := os.Open(logFile); err != nil {
	//	t.Error(err)
	//	t.FailNow()
	//}
}
