package godebug

import "testing"

func TestPrint(t *testing.T) {
	Info("debug test message")
	Infof("%s", "debug test message")
}
