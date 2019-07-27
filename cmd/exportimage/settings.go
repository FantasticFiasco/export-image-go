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
func newSettings() (s settings, err error) {
    exec, err := os.Executable()
    if err != nil {
        return
    }

	f := filepath.Join(filepath.Dir(exec), "exportimage.yml")

	data, err := ioutil.ReadFile(f)
    if err != nil {
        return
    }

	err = yaml.Unmarshal(data, &s)
    if err != nil {
        return
    }

	return
}
