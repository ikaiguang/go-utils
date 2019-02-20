package goerror

import (
	"fmt"
	"runtime"
)

// error code
const (
	OK       = 0 // success code
	Canceled = 1 // cancel code
	Unknown  = 2 // unknown code
)

// Status : defines a logical error model
type Status struct {
	Code    int    `json:"code"`    // HTTP status codes
	Message string `json:"message"` // message
	Caller  string `json:"-"`       // caller
}

// Error : error implementation
func (s *Status) Error() string {
	return fmt.Sprintf("code(%d) %s", s.Code, s.Message)
}

// Format Formatter
func (s *Status) Format(fmtState fmt.State, verb rune) {
	fmtState.Write([]byte(s.Detail()))
}

// Detail error detail inof
func (s *Status) Detail() string {
	printFormat := `
~~~~~ ~~~~~ ~~~~~
error :
	code : %d
	message : %s
	caller : %s
`
	return fmt.Sprintf(printFormat, s.Code, s.Message, s.Caller)
}

// New returns a Status representing c and msg.
func New(code int, msg string, customCallerSkip ...int) error {
	return newStatus(code, msg, customCallerSkip...)
}

// newStatus new Status
func newStatus(code int, msg string, customCallerSkip ...int) *Status {
	// call skip
	var skip = 1
	if len(customCallerSkip) > 0 {
		skip = customCallerSkip[0]
	}

	// caller
	pc, file, line, _ := runtime.Caller(skip + 1)

	return &Status{
		Code:    code,
		Message: msg,
		Caller:  callerFormat(pc, file, line),
	}
}

// callerFormat call format
func callerFormat(pc uintptr, file string, line int) string {
	return fmt.Sprintf("%s \n\t\t %s:%d", runtime.FuncForPC(pc).Name(), file, line)
}

// NewWithError error over error
func NewWithError(code int, msg string, err error, customCallerSkip ...int) error {
	newStatus := newStatus(code, msg, customCallerSkip...)

	if s, ok := FromError(err); ok {
		newStatus.Caller += s.Detail()
	} else {
		newStatus.Caller += stdErrFormat(err)
	}
	return newStatus
}

// stdErrFormat standard error format
func stdErrFormat(err error) string {
	return "\n~~~~~ ~~~~~ ~~~~~\nerror : \n\t" + err.Error()
}

// FromError returns a Status representing err if it was produced from this package,
// otherwise it returns nil, false.
func FromError(err error) (*Status, bool) {
	if err == nil {
		return &Status{Code: OK}, true
	}
	if status, ok := err.(*Status); ok {
		return status, true
	}
	return nil, false
}

// Forward forward error
func Forward(err error) error {
	newStatus := newStatus(Unknown, err.Error())

	if s, ok := FromError(err); ok {
		newStatus.Code = s.Code
		newStatus.Caller += s.Detail()
	} else {
		newStatus.Caller += stdErrFormat(err)
	}
	return newStatus
}
