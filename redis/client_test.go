package goredis

import (
	"os"
	"testing"
)

func TestNewClient(t *testing.T) {
	os.Setenv(EnvKeyRedisAddress, "127.0.0.1:6379")
	os.Setenv(EnvKeyRedisPassword, "")
	os.Setenv(EnvKeyRedisDB, "0")
	os.Setenv(EnvKeyRedisPoolSize, "0")
	os.Setenv(EnvKeyRedisMinIdleConns, "0")
	os.Setenv(EnvKeyRedisMaxConnAge, "0s")
	os.Setenv(EnvKeyRedisDialTimeout, "0s")
	os.Setenv(EnvKeyRedisReadTimeout, "0s")
	os.Setenv(EnvKeyRedisWriteTimeout, "0s")
	os.Setenv(EnvKeyRedisPoolTimeout, "0s")
	os.Setenv(EnvKeyRedisIdleTimeout, "0s")
	os.Setenv(EnvKeyRedisIdleCheckFrequency, "0s")
	os.Setenv(EnvKeyRedisMaxRetries, "0")
	os.Setenv(EnvKeyRedisMinRetryBackoff, "0s")
	os.Setenv(EnvKeyRedisMaxRetryBackoff, "0s")

	if _, err := NewClient(); err != nil {
		t.Error(err)
		t.FailNow()
	}
}
