package goerror

import (
	"io"
	"testing"
)

func TestNew(t *testing.T) {
	err := New(1, "warm prompt")

	t.Log(err.Error())    // {"code":1,"message":"warm prompt"}
	t.Logf("%v \n", err)  // {"code":1,"message":"warm prompt"}
	t.Logf("%#v \n", err) // {"code":1,"message":"warm prompt"}
	t.Logf("%+v \n", err) // {"code":1,"message":"warm prompt"} + \n\t + funcName + \n\t\t + file:line
}

func TestWarp(t *testing.T) {
	err := Warp(io.EOF, 1, "warm prompt")
	t.Log(err.Error())    // EOF: {"code":1,"message":"warm prompt"}
	t.Logf("%#v \n", err) // EOF: {"code":1,"message":"warm prompt"}
	t.Logf("%+v \n", err) // {"code":1,"message":"warm prompt"} + \n + EOF + \n\t + funcName + ...
}

func TestForward(t *testing.T) {
	err := New(1, "warm prompt")
	t.Log(err)
	t.Logf("%+v \n", err) // {"code":1,"message":"warm prompt"} + \n\t + funcName + \n\t\t + file:line

	err = Forward(err)
	t.Log(err)            // {"code":1,"message":"warm prompt"}
	t.Logf("%+v \n", err) // {"code":1,"message":"warm prompt"} + \n\t + funcName + \n\t\t + file:line
}

func TestFromError(t *testing.T) {
	f, ok := FromError(New(1, "warm prompt"))
	if !ok {
		t.FailNow()
	}
	t.Log(f.Error()) // {"code":1,"message":"warm prompt"}
}

func TestForwardAndWarp(t *testing.T) {
	err := WithStack(io.EOF)
	err = Forward(err)
	err = Forward(err)
	t.Log(err)            // EOF
	t.Logf("%+v \n", err) // EOF + \n\t + funcName + ...
	err = Warp(err, 1, "warm prompt")
	t.Log(err)    // {"code":1,"message":"warm prompt"}
	t.Logf("%+v \n", err) // {"code":1,"message":"warm prompt"} + \n + EOF + \n\t + funcName + ...
	err = Warp(err, 1, "error")
	t.Log(err.Error())
	t.Logf("%+v \n", err)
	err = Warp(err, 1, "error")
	t.Log(err.Error())
	t.Logf("%+v \n", err)
	f, ok := FromError(err)
	if !ok {
		t.FailNow()
	}
	t.Log(f.Error()) // {"code":1,"message":"warm prompt"}
}
