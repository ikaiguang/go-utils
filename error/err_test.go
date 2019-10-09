package goerror

import (
	"testing"
)

func TestNew(t *testing.T) {
	// max callers
	SetMaxCallers(1)

	// new
	err := Err(1, "this is test error")
	t.Log(err)
	t.Logf("%+v \n", err)

	// wrap
	err = Wrap(err, "this is wrap xxx")
	t.Logf("%+v \n", err)

	// from
	if s, ok := FromErr(err); !ok {
		t.Error("is test error : ", ok)
	} else {
		t.Log("is test error : ", ok)
		t.Log(s)
	}
}
