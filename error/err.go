package goerror

import (
	stderrors "errors"
	"fmt"
	"io"
)

// status code
const (
	OK       = 0 // status ok
	Canceled = 1 // status cancel
	Unknown  = 2 // status unknown
)

// New returns an error from Status
func New(code int32, message string) error {
	return &Status{
		Code:    code,
		Message: message,
		stack:   callers(),
	}
}

// NewWithDetails returns an error from Status
func NewWithDetails(code int32, message string, details interface{}) error {
	return &Status{
		Code:    code,
		Message: message,
		Details: details,
		stack:   callers(),
	}
}

// Newf returns an error from Status
func Newf(code int32, format string, args ...interface{}) error {
	return &Status{
		Code:    code,
		Message: fmt.Sprintf(format, args...),
		stack:   callers(),
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

// FromError returns a Status representing err if it was produced from this package.
// Otherwise, ok is false and a Status is returned with code(Unknown) and the original error message.
func FromError(err error) (s *Status, ok bool) {
	if err == nil {
		return &Status{Code: OK, Message: "", stack: callers()}, true
	}
	if s, ok := Cause(err).(*Status); ok {
		return s, true
	}
	return &Status{Code: Unknown, Message: err.Error(), stack: callers()}, false
}

// Status implements error
type Status struct {
	Code    int32       `json:"code"`              // code
	Message string      `json:"message"`           // message
	Details interface{} `json:"details,omitempty"` // details
	*stack  `json:"-"`                             // stack
}

// Error implements error
func (s *Status) Error() string {
	return fmt.Sprintf("go error: code = %d msg = %s", s.Code, s.Message)
}

// ErrorWithDetail error with detail
func (s *Status) ErrorWithDetail() string {
	return fmt.Sprintf("go error: code = %d meg = %s \n \t details = %v", s.Code, s.Message, s.Details)
}

// ErrDetail error detail
func (s *Status) ErrDetail() string {
	return fmt.Sprintf("%v", s.Details)
}

// Format format
func (s *Status) Format(state fmt.State, verb rune) {
	switch verb {
	case 'v':
		if state.Flag('+') {
			io.WriteString(state, s.ErrorWithDetail())
			s.stack.Format(state, verb)
			return
		}
		fallthrough
	case 's':
		io.WriteString(state, s.ErrorWithDetail())
	case 'q':
		fmt.Fprintf(state, "%q", s.Error())
	}
}

type withStack struct {
	error
	*stack
}

func (w *withStack) Cause() error { return w.error }

// Unwrap provides compatibility for Go 1.13 error chains.
func (w *withStack) Unwrap() error { return w.error }

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

// Cause recursively unwraps an error chain and returns the underlying cause of
// the error, if possible. There are two ways that an error value may provide a
// cause. First, the error may implement the following interface:
//
//     type causer interface {
//            Cause() error
//     }
//
// Second, the error may return a non-nil value when passed as an argument to
// the Unwrap function. This makes Cause forwards-compatible with Go 1.13 error
// chains.
//
// If an error value satisfies both methods of unwrapping, Cause will use the
// causer interface.
//
// If the error is nil, nil will be returned without further investigation.
func Cause(err error) error {
	type causer interface {
		Cause() error
	}

	for err != nil {
		if cause, ok := err.(causer); ok {
			err = cause.Cause()
		} else if unwrapped := Unwrap(err); unwrapped != nil {
			err = unwrapped
		} else {
			break
		}
	}
	return err
}

// Unwrap returns the result of calling the Unwrap method on err, if err's
// type contains an Unwrap method returning error.
// Otherwise, Unwrap returns nil.
func Unwrap(err error) error {
	return stderrors.Unwrap(err)
}
