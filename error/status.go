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
	Caller  string `json:"caller"`  // caller
}

// Error : error implementation
func (s *Status) Error() string {
	return fmt.Sprintf("code :%d \n message : %s", s.Code, s.Message)
}

// Format Formatter
func (s *Status) Format(fmtState fmt.State, verb rune) {
	fmt.Println(s.Detail())
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
		Caller:  fmt.Sprintf("%s \n\t\t %s:%d", runtime.FuncForPC(pc).Name(), file, line),
	}
}

// NewWithError error over error
func NewWithError(code int, msg string, err error, customCallerSkip ...int) error {
	newStatus := newStatus(code, msg, customCallerSkip...)

	errStatus, errStatusOK := FromError(err)
	if errStatusOK {
		newStatus.Caller += errStatus.Detail()
	} else {
		newStatus.Caller += "\n~~~~~ ~~~~~ ~~~~~\nerror : \n\t" + err.Error()
	}
	return newStatus
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
