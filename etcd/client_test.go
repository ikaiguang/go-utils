package goetcd

import (
	"os"
	"testing"
)

func TestNewClient(t *testing.T) {
	os.Setenv(EnvKeyETCDEndPoints, "127.0.0.1:2379,127.0.0.1:2379")
	os.Setenv(EnvKeyETCDUsername, "")
	os.Setenv(EnvKeyETCDPassword, "")
	os.Setenv(EnvKeyETCDDialTimeout, "3s")
	os.Setenv(EnvKeyETCDDialKeepAliveTime, "0s")
	os.Setenv(EnvKeyETCDDialKeepAliveTimeout, "0s")
	os.Setenv(EnvKeyETCDAutoSyncInterval, "0s")
	os.Setenv(EnvKeyETCDMaxCallSendMsgSize, "0")
	os.Setenv(EnvKeyETCDMaxCallRecvMsgSize, "0")
	os.Setenv(EnvKeyETCDRejectOldCluster, "false")

	if _, err := NewClient(); err != nil {
		t.Error(err)
		t.FailNow()
	}
}
