package goruntime

import "testing"

func TestFile(t *testing.T) {
	t.Log(File())
}

func TestLine(t *testing.T) {
	t.Log(Line())
}

func TestFunc(t *testing.T) {
	t.Log(Func())
}
