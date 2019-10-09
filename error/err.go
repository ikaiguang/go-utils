package goerror

import (
	"fmt"
	"io"
)

// Err returns an error from Status
func Err(code int32, message string) error {
	return &Status{
		Code:    code,
		Message: message,
		stack:   callers(),
	}
}

// Errf returns an error from Status
func Errf(code int32, format string, args ...interface{}) error {
	return &Status{
		Code:    code,
		Message: fmt.Sprintf(format, args...),
		stack:   callers(),
	}
}

// FromErr returns a Status representing err if it was produced from this package.
// Otherwise, ok is false and a Status is returned with code(0) and the original error message.
func FromErr(err error) (s *Status, ok bool) {
	if err == nil {
		return &Status{Code: 0, Message: err.Error(), stack: callers()}, true
	}
	if s, ok := Cause(err).(*Status); ok {
		return s, true
	}
	return &Status{Code: 0, Message: err.Error(), stack: callers()}, false
}

// Status error
type Status struct {
	Code    int32  `json:"code"`    // code
	Message string `json:"message"` // message
	*stack  `json:"-"`              // stack
}

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
