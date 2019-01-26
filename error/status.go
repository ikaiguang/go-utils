package goerror

import (
	"fmt"
	"net/http"
	"runtime"
)

// OK : http StatusOK(package net/http)
// HTTP status codes as registered with IANA.
// See: https://www.iana.org/assignments/http-status-codes/http-status-codes.xhtml
const OK = http.StatusOK

// Status : defines a logical error model
type Status struct {
	Code    int    `json:"code"`    // HTTP status codes
	Message string `json:"message"` // message
	Caller  string `json:"caller"`  // caller
}

// Error : error implementation
func (s *Status) Error() string {
	printFormat := `
~~~~~ ~~~~~ ~~~~~
error :
	code : %d
	message : %s
	caller : %s`

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

	oldStatus, ok := FromError(err)
	if ok {
		newStatus.Caller += oldStatus.Error()
	} else{
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
