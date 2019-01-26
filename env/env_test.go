package goenv

import "testing"

func TestGetEnv(t *testing.T) {

	t.Log(GetEnv("GOPATH"))

	t.Log(GetEnv("abc123"))
}
