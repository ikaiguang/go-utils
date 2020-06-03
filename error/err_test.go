package goerror

import (
	"github.com/pkg/errors"
	"io"
	"testing"
)

func TestNew(t *testing.T) {
	// process error
	var eofFn = func() error { return errors.WithStack(io.EOF) }
	var myFn = func() error { return New(OK, "anything is ok") }
	var detailsFn = func() error { return NewWithDetails(OK, "anything is ok", "test details") }

	// eof error
	eofErr := eofFn()
	t.Logf("eof error : %+v \n", eofErr)
	eofStackErr := errors.WithStack(eofErr)
	// cannot parse
	if _, ok := FromError(eofStackErr); ok {
		t.Error("something write wrong")
		return
	}

	// my error
	myErr := myFn()
	t.Logf("my error : %+v \n", myErr)
	myStackErr := errors.WithStack(myErr)
	t.Logf("my stack error : %+v \n", myStackErr)
	// parse error
	myStatus, ok := FromError(myStackErr)
	if !ok {
		t.Error("something write wrong")
		return
	}
	t.Logf("my status : code : %d, message : %s \n", myStatus.Code, myStatus.Message)

	// details error
	dErr := detailsFn()
	t.Logf("details error : %+v \n", dErr)
	dStackErr := WithStack(dErr)
	dStackErr = WithStack(dStackErr)
	t.Logf("details stack error : %+v \n", dStackErr)
	// parse error
	dStatus, ok := FromError(dStackErr)
	if !ok {
		t.Error("something write wrong")
		return
	}
	t.Logf(dStatus.ErrorWithDetail())
}
