package gocfg

import (
	"path/filepath"
	"runtime"
)

// currentPath is the root directory of this package.
var currentPath string

func init() {
	_, currentFile, _, _ := runtime.Caller(0)
	currentPath = filepath.Dir(currentFile)
}

// Path returns the absolute path the given relative file or directory path
func Path(rel string) string {
	if filepath.IsAbs(rel) {
		return rel
	}
	return filepath.Join(currentPath, rel)
}
