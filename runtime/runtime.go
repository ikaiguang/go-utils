package goruntime

import (
	"runtime"
)

// File returns the file name in which the function was invoked
func File() string {
	_, file, _, _ := runtime.Caller(1)
	return file
}

// Line returns the line number at which the function was invoked
func Line() int {
	_, _, line, _ := runtime.Caller(1)
	return line
}

// Func returns the name of the function.
func Func() string {
	pc, _, _, _ := runtime.Caller(1)

	return runtime.FuncForPC(pc).Name()
}
