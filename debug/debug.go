package godebug

import (
	"fmt"
	"runtime"
)

// closeOutput close output
var closeOutput bool

// CloseOutput close output
func CloseOutput() {
	closeOutput = true
}

// Info info
func Info(a ...interface{}) {
	if closeOutput {
		return
	}
	caller()
	fmt.Println(a...)
}

// Infof info
func Infof(format string, a ...interface{}) {
	if closeOutput {
		return
	}
	caller()
	format += "\n"
	fmt.Printf(format, a...)
}

// caller call
func caller() {
	_, file, line, _ := runtime.Caller(2)
	fmt.Printf("==> debug file : %s(%d) \n", file, line)
}

// emptyWriter empty writer
type emptyWriter struct{}

// Write writer
func (e *emptyWriter) Write(p []byte) (n int, err error) {
	return
}
