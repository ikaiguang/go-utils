package goerror

import (
	"fmt"
	"github.com/pkg/errors"
	"io"
)

// status code
const (
	OK      = 0  // status ok
	Unknown = -1 // status unknown
)

// New returns an error from Status
func New(code int32, message string) error {
	return &Status{
		Code:    code,
		Message: message,
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

// FromError returns a Status representing err if it was produced from this package.
// Otherwise, ok is false and a Status is returned with code(Unknown) and the original error message.
func FromError(err error) (s *Status, ok bool) {
	if err == nil {
		return &Status{Code: OK, Message: "", stack: callers()}, true
	}
	if s, ok := errors.Cause(err).(*Status); ok {
		return s, true
	}
	return &Status{Code: Unknown, Message: err.Error(), stack: callers()}, false
}

// Status implements error
type Status struct {
	Code    int32  `json:"code"`    // code
	Message string `json:"message"` // message
	*stack  `json:"-"`              // stack
}

// Error implements error
func (s *Status) Error() string {
	return fmt.Sprintf("go error: code = %d desc = %s", s.Code, s.Message)
}

// Format format
func (s *Status) Format(state fmt.State, verb rune) {
	switch verb {
	case 'v':
		if state.Flag('+') {
			io.WriteString(state, s.Error())
			s.stack.Format(state, verb)
			return
		}
		fallthrough
	case 's':
		io.WriteString(state, s.Error())
	case 'q':
		fmt.Fprintf(state, "%q", s.Error())
	}
}
