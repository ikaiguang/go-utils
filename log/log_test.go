package golog

import (
	"os"
	"testing"
	"time"
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

func TestNewLog_Rotation_3s(t *testing.T) {
	os.Setenv(EnvKeyLogMysqlEnable, "false")
	os.Setenv(EnvKeyLogFileEnable, "true")
	os.Setenv(EnvKeyLogFileRotation, "true")
	os.Setenv(EnvKeyLogFilename, "_output/10s_log")
	os.Setenv(EnvKeyLogFileOptTimeLocation, "Local")
	os.Setenv(EnvKeyLogFileOptLinkName, "_output/10s_log")
	os.Setenv(EnvKeyLogFileOptForceNewFile, "false")
	os.Setenv(EnvKeyLogFileOptMaxAge, "8760h")
	os.Setenv(EnvKeyLogFileOptMaxAge, "0s")
	os.Setenv(EnvKeyLogFileOptRotationTime, "3s")
	//os.Setenv(EnvKeyLogFileOptRotationCount, "10000")
	os.Setenv(EnvKeyLogFileOptRotationSuffix, ".%Y_%m_%d_%H_%M_%S")

	handler, err := NewLog()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	ticker := time.NewTicker(time.Second)
	for {
		select {
		case <-ticker.C:
			handler.Info("info message, time : ", time.Now())
		}
	}

	//logFile := "./test.log." + time.Now().Format("2006_01_02")
	//if _, err := os.Open(logFile); err != nil {
	//	t.Error(err)
	//	t.FailNow()
	//}
}
