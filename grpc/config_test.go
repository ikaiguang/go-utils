package gogrpc

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
	"runtime"
	"testing"
)

func TestParseConfig(t *testing.T) {
	filename := "config.yml"

	_, currentFile, _, _ := runtime.Caller(0)
	currentPath := filepath.Dir(currentFile)

	fullFilename := filepath.Join(currentPath, filename)

	// read file
	fileBytes, err := ioutil.ReadFile(fullFilename)
	if err != nil {
		t.Error(err)
		return
	}
	_ = fileBytes

	var cfg = struct {
		Config `yaml:"gogrpc"`
	}{}
	if err := yaml.Unmarshal(fileBytes, &cfg); err != nil {
		t.Error(err)
		return
	}
	t.Logf("%#v \n",cfg.Config)
}
