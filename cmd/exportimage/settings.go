package main

import (
	"io/ioutil"
	"os"
	"path/filepath"

    "github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

// Settings is wrapping the YAML configuration file.
type settings struct {
	Quality            int    `yaml:"quality"`
	SubDirectory       string `yaml:"subDirectory"`
	PreventTermination bool   `yaml:"preventTermination"`
}

// New allocates and returns a new Encoder.
func newSettings() (settings, error) {
    var s settings

    exec, err := os.Executable()
    if err != nil {
        return s, errors.Wrap(err, "get running executable failed")
    }

	f := filepath.Join(filepath.Dir(exec), "exportimage.yml")

	data, err := ioutil.ReadFile(f)
    if err != nil {
        return s, errors.Wrap(err, "read settings file failed")
    }

	err = yaml.Unmarshal(data, &s)
    if err != nil {
        return  s, errors.Wrap(err, "deserializing settings file failed")
    }

	return s, nil
}
