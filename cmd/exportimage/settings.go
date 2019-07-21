package main

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

// Settings is wrapping the YAML configuration file.
type settings struct {
	Quality            int    `yaml:"quality"`
	SubDirectory       string `yaml:"subDirectory"`
	PreventTermination bool   `yaml:"preventTermination"`
}

// New allocates and returns a new Encoder.
func newSettings() settings {
	exec, err := os.Executable()
	check(err)

	f := filepath.Join(filepath.Dir(exec), "exportimage.yml")

	data, err := ioutil.ReadFile(f)
	check(err)

	var s settings
	err = yaml.Unmarshal(data, &s)
	check(err)

	return s
}
