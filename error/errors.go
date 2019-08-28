// from github.com/pkg/errors
package goerror

import (
	"encoding/json"
	"fmt"
	"io"
)

// FromError returns a Status representing err if it was produced from this package,
// otherwise it returns nil, false.
func FromError(err error) (*Status, bool) {
	if f, ok := Cause(err).(*fundamental); ok {
		return f.status, ok
	}
	return nil, false
}

// New returns an error with the supplied message.
// New also records the stack trace at the point it was called.
func New(code int, message string) error {
	return &fundamental{
		status: &Status{Code: code, Msg: message},
		stack:  callers(),
	}
}

type Status struct {
	Code int    `json:"code"`
	Msg  string `json:"message"`
}

func (s *Status) Error() string { b, _ := json.Marshal(s); return string(b) }

// fundamental is an error that has a message and a stack, but no caller.
type fundamental struct {
	status *Status
	*stack
}

func (f *fundamental) Error() string { return f.status.Error() }
func (f *fundamental) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		if s.Flag('+') {
			io.WriteString(s, f.Error())
			f.stack.Format(s, verb)
			return
		}
		fallthrough
	case 's':
		io.WriteString(s, f.Error())
	case 'q':
		fmt.Fprintf(s, "%q", f.Error())
	}
}

// Forward annotates err with a stack trace at the point WithStack was called.
// If err is nil, WithStack returns nil.
func Forward(err error) error {
	if err == nil {
		return nil
	}
	return &withStack{
		err,
		callers(),
	}
}

// WithStack annotates err with a stack trace at the point WithStack was called.
// If err is nil, WithStack returns nil.
func WithStack(err error) error {
	if err == nil {
		return nil
	}
	return &withStack{
		err,
		callers(),
	}
}

type withStack struct {
	error
	*stack
}

func (w *withStack) Cause() error { return w.error }
func (w *withStack) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		if s.Flag('+') {
			fmt.Fprintf(s, "%+v", w.Cause())
			w.stack.Format(s, verb)
			return
		}
		fallthrough
	case 's':
		io.WriteString(s, w.Error())
	case 'q':
		fmt.Fprintf(s, "%q", w.Error())
	}
}

// Wrap returns an error annotating err with a stack trace
// at the point Wrap is called, and the supplied message.
// If err is nil, Wrap returns nil.
// if Cause(err).(*fundamental) is ok, forward err with stack
func Wrap(err error, code int, message string) error {
	if err == nil {
		return nil
	}
	// forward
	if _, ok := Cause(err).(*fundamental); ok {
		return &withStack{
			err,
			callers(),
		}
	}
	// new error
	err = &withMessage{
		cause: &fundamental{
			status: &Status{Code: code, Msg: message},
			stack:  callers(),
		},
		err: err,
		//msg: err.Error(),
	}
	return &withStack{
		err,
		callers(),
	}
}

type withMessage struct {
	cause error
	err   error
	//msg   string
}

func (w *withMessage) Error() string { return w.cause.Error() }
func (w *withMessage) Cause() error  { return w.cause }
func (w *withMessage) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		if s.Flag('+') {
			fmt.Fprintf(s, "%+v\n", w.Cause())
			fmt.Fprintf(s, "%+v\n", w.err)
			//io.WriteString(s, w.msg)
			return
		}
		fallthrough
	case 's', 'q':
		io.WriteString(s, w.Error())
	}
}

// Cause returns the underlying cause of the error, if possible.
// An error value has a cause if it implements the following
// interface:
//
//     type causer interface {
//            Cause() error
//     }
//
// If the error does not implement Cause, the original error will
// be returned. If the error is nil, nil will be returned without further
// investigation.
func Cause(err error) error {
	type causer interface {
		Cause() error
	}

	for err != nil {
		cause, ok := err.(causer)
		if !ok {
			break
		}
		err = cause.Cause()
	}
	return err
}
