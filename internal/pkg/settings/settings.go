package settings

import (
    "github.com/FantasticFiasco/export-image-go/internal/pkg/errors"
    "gopkg.in/yaml.v2"
    "io/ioutil"
    "os"
    "path/filepath"
)

// Settings is wrapping the YAML configuration file.
type Settings struct {
	Quality int `yaml:"quality"`
}

// New allocates and returns a new Encoder.
func New() Settings {
    exec, err := os.Executable()
    errors.Check(err)

    file := filepath.Join(filepath.Dir(exec), "exportimage.yml")

    data, err := ioutil.ReadFile(file)
	errors.Check(err)

    s := Settings{}
	err = yaml.Unmarshal(data, &s)
	errors.Check(err)

	return s
}
